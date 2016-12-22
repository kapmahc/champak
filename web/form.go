package web

import (
	"github.com/gorilla/csrf"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Form form
type Form map[string]interface{}

// AddFields add fields
func (p Form) AddFields(fields ...interface{}) {
	items := p["fields"].([]interface{})
	items = append(items, fields...)
	p["fields"] = items
}

// NewForm new form
func NewForm(ctx *gin.Context, id, title, action string) Form {
	fm := make(Form)
	fm["locale"] = ctx.MustGet(LOCALE)
	fm[csrf.TemplateTag] = csrf.TemplateField(ctx.Request)
	fm["title"] = title
	fm["method"] = "post"
	fm["action"] = action
	fm["fields"] = make([]interface{}, 0)
	return fm
}

// TextField text field
type TextField struct {
	Require     bool
	Type        string
	ID          string
	Label       string
	Value       string
	Help        string
	Placeholder string
	ReadOnly    bool
}

// NewTextField new text field
func NewTextField(id, label, value string) *TextField {
	return &TextField{
		Require: true,
		Type:    "text",
		ID:      id,
		Label:   label,
		Value:   value,
	}
}

// EmailField email field
type EmailField struct {
	Require     bool
	ReadOnly    bool
	Type        string
	ID          string
	Label       string
	Value       string
	Help        string
	Placeholder string
}

// NewEmailField new email field
func NewEmailField(id, label, value string) *EmailField {
	return &EmailField{
		Require: true,
		Type:    "email",
		ID:      id,
		Label:   label,
		Value:   value,
	}
}

// PasswordField password field
type PasswordField struct {
	Require     bool
	Type        string
	ID          string
	Label       string
	Help        string
	Placeholder string
}

// NewPasswordField new password field
func NewPasswordField(id, label string) *PasswordField {
	return &PasswordField{
		Require: true,
		Type:    "password",
		ID:      id,
		Label:   label,
	}
}

// HiddenField hidden field
type HiddenField struct {
	Type  string
	ID    string
	Value interface{}
}

// NewHiddenField new hidden field
func NewHiddenField(id string, value interface{}) *HiddenField {
	return &HiddenField{
		Type:  "hidden",
		ID:    id,
		Value: value,
	}
}
