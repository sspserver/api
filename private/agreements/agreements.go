package agreements

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

//go:embed *.md
var FS embed.FS

var (
	mdParser = goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}))
	validate = validator.New(validator.WithRequiredStructEnabled())
)

type AgreementMeta struct {
	Codename    string    `json:"codename" toml:"codename" validate:"required"`
	Title       string    `json:"title" toml:"title" validate:"required"`
	Description string    `json:"description" toml:"description"`
	Version     string    `json:"version" toml:"version" validate:"required"`
	Type        string    `json:"type" toml:"type" validate:"oneof=license terms_of_use contract"` // e.g. "license", "terms_of_use", "contract"
	IssuedBy    string    `json:"issuedBy" toml:"issuedBy" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" toml:"createdAt" validate:"required"` // ISO 8601 format
}

type Agreement struct {
	Meta         AgreementMeta `json:"meta"`
	Body         string        `json:"body"`
	BodyHTML     string        `json:"bodyHtml"`
	BodyMarkdown string        `json:"bodyMarkdown"`
}

func Agreements() []*Agreement {
	agreements := make([]*Agreement, 0)
	_ = fs.WalkDir(FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := fs.ReadFile(FS, path)
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".md") {
			mdData, htmlData, meta, err := renderMarkdown(data)
			if err != nil {
				panic("renderMarkdown: " + err.Error())
			}
			agreements = append(agreements, &Agreement{
				Meta:         *meta,
				Body:         string(data),
				BodyHTML:     string(htmlData),
				BodyMarkdown: string(mdData),
			})
		}
		return nil
	})
	return agreements
}

func renderMarkdown(data []byte) (md, html string, meta *AgreementMeta, err error) {
	buf := bytes.Buffer{}
	ctx := parser.NewContext()

	if err = mdParser.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return "", "", nil, err
	}

	// Remove leading and trailing whitespace from the data
	data = bytes.TrimSpace(data)

	// Extract frontmatter
	meta = &AgreementMeta{}
	if data := frontmatter.Get(ctx); data != nil {
		if err := data.Decode(meta); err != nil {
			return "", "", nil, fmt.Errorf("decode frontmatter: %w", err)
		}
	}

	if err := validate.Struct(meta); err != nil {
		return "", "", nil, fmt.Errorf("validate meta: %w", err)
	}

	// Remove "+++" and "---" from the beginning and end of the body
	if len(data) >= 3 {
		blockType := string(data[:3])
		if blockType == "---" || blockType == "+++" {
			data = data[3:]

			// Find the end of the frontmatter
			endIndex := bytes.Index(data, []byte("\n"+blockType+"\n"))
			if endIndex != -1 {
				data = data[endIndex+5:] // Skip the end marker and newline
			}
		}
	}

	return string(data), buf.String(), meta, nil
}
