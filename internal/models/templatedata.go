package models

import (
	"github.com/fakorede/gobnb/internal/forms"
)

// TemplateData holds possible data sent to templates from handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
