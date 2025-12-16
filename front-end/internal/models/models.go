package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]any
}

type Userquestion struct {
	Question string `json:"question"`
}

type AgentResponse struct {
	Answer string `json:"answer"`
}

type JSONResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
