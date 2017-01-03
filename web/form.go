package web

import (
	"strconv"

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
	fm["next"] = ctx.Request.URL.Path
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

// TextArea textarea
type TextArea struct {
	Require     bool
	Type        string
	ID          string
	Label       string
	Value       string
	Help        string
	Placeholder string
	ReadOnly    bool
}

// NewTextArea new textarea
func NewTextArea(id, label, value string) *TextArea {
	return &TextArea{
		Require: true,
		Type:    "textarea",
		ID:      id,
		Label:   label,
		Value:   value,
	}
}

// Select select
type Select struct {
	Require     bool
	Type        string
	ID          string
	Label       string
	Help        string
	Placeholder string
	ReadOnly    bool
	Multiple    bool
	Options     []Option
}

// Option option
type Option struct {
	Label    string
	Value    interface{}
	Selected bool
}

// NewOrderSelect new sort order  select
func NewOrderSelect(id, label string, value, min, max int) *Select {
	var options []Option
	for i := min; i <= max; i++ {
		options = append(options, Option{Label: strconv.Itoa(i), Value: value, Selected: i == value})
	}

	return NewSelect(id, label, options)
}

// NewSelect new select
func NewSelect(id, label string, options []Option) *Select {
	return &Select{
		Require: true,
		Type:    "select",
		ID:      id,
		Label:   label,
		Options: options,
	}
}

// Checkbox checkbox
type Checkbox struct {
	Require     bool
	Type        string
	ID          string
	Label       string
	Value       bool
	Help        string
	Placeholder string
	ReadOnly    bool
}

// NewCheckbox new select
func NewCheckbox(id, label string, value bool) *Checkbox {
	return &Checkbox{
		Require: true,
		Type:    "checkbox",
		ID:      id,
		Label:   label,
		Value:   value,
	}
}
