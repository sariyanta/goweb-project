package models

type TemplateData struct {
	StringMap  map[string]string
	Navigation map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CSRFToken  string
	Flash      string
	Warning    string
	Form       string
	Error      string
}
