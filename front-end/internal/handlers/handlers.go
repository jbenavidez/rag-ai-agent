package handlers

import (
	"fmt"
	"frontend/internal/config"
	"frontend/internal/models"
	"frontend/internal/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// ChatRoom is the handler for the chatroom page
func (m *Repository) ChatRoom(w http.ResponseWriter, r *http.Request) {

	// Render template
	render.RenderTemplate(w, r, "chatroom.page.tmpl", &models.TemplateData{
		Data: nil,
	})
}

// AnswerUserQuestion: answer user question.
func (m *Repository) AnswerUserQuestion(w http.ResponseWriter, r *http.Request) {
	var payload models.Userquestion

	err := m.readJSON(w, r, &payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("the user question", payload)
	// set response
	agentAnwer := models.AgentResponse{
		Answer: "hello from gondor",
	}
	//set response
	resp := models.JSONResponse{
		Error:   false,
		Message: "answer for qeustion",
		Data:    agentAnwer,
	}
	// send resposne
	fmt.Println("the response", resp)
	_ = m.writeJSON(w, http.StatusOK, resp)

}
