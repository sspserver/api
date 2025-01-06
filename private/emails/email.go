package emails

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	htmltemplate "html/template"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter" // https://github.com/abhinav/goldmark-frontmatter

	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/sendmsg/template"
	"github.com/demdxx/sendmsg/template/godef"
	"github.com/demdxx/xtypes"
)

//go:embed *.md
var FS embed.FS

var mdParser = goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}))

func Templates() template.Storage {
	store := template.NewDefaultStorage()
	_ = fs.WalkDir(FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := fs.ReadFile(FS, path)
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".md") {
			name := strings.TrimSuffix(filepath.Base(path), ".md")
			nData, meta, err := renderMarkdown(data)
			if err != nil {
				panic("renderMarkdown: " + err.Error())
			}

			tmpl, err := godef.NewHTMLTemplate(name,
				gocast.Str(meta["subject"]),
				string(nData),
				godef.WithHTMLVars(gocast.Map[string, any](meta)),
				godef.WithHTMLPreRender(emailPrerender),
			)
			if err != nil {
				panic("godef.NewHTMLTemplate: " + err.Error())
			}
			store.RegisterTmpl(tmpl)
		}
		return nil
	})
	return store
}

func emailPrerender(ctx context.Context, vars map[string]any) (map[string]any, error) {
	if urlFormat, ok := vars["urlPasswordResetFormat"]; ok {
		varCopy := xtypes.Map[string, any](vars).Copy()
		passResetURL := strings.NewReplacer(
			"{{email}}", gocast.Str(vars["email"]),
			"{{code}}", gocast.Str(vars["reset_token"]),
		).Replace(gocast.Str(urlFormat))
		varCopy["urlPasswordReset"] = htmltemplate.HTML(passResetURL)
		return varCopy, nil
	}
	return vars, nil
}

func renderMarkdown(data []byte) ([]byte, map[string]string, error) {
	buf := bytes.Buffer{}
	ctx := parser.NewContext()

	if err := mdParser.Convert(data, &buf, parser.WithContext(ctx)); err != nil {
		return nil, nil, err
	}

	var meta map[string]string
	if data := frontmatter.Get(ctx); data != nil {
		if err := data.Decode(&meta); err != nil {
			return nil, nil, fmt.Errorf("decode frontmatter: %w", err)
		}
	}

	return buf.Bytes(), meta, nil
}
