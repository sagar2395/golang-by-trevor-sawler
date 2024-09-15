package models

// template data sends data to handlers
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int32
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
