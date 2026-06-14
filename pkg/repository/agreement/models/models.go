package models

import "time"

// Agreement represents a legal or contractual agreement
type Agreement struct {
	Codename    string `json:"codename"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Type        string `json:"type"` // e.g. "license", "terms_of_use", "contract"
	IssuedBy    string `json:"issued_by"`

	BodyMarkdown string `json:"body_md"`
	BodyHTML     string `json:"body_html"`

	AcceptAccountID uint64     `json:"accept_account_id"`   // ID of the account that accepted the agreement
	AcceptByUserID  uint64     `json:"accept_by_user_id"`   // ID of the user that accepted the agreement
	Signature       string     `json:"signature,omitempty"` // Base64 encoded signature, optional
	AcceptedAt      *time.Time `json:"accepted_at"`         // ISO 8601 format

	CreatedAt time.Time `json:"createdAt"` // ISO 8601 format
}

// RBACResourceName returns the resource name for RBAC.
func (a *Agreement) RBACResourceName() string {
	return "agreement"
}
