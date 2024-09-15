package handlers

import (
	"fmt"
	"net/http"

	"github.com/sagar2395/golang-by-trevor-sawler/pkg/config"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/models"
	"github.com/sagar2395/golang-by-trevor-sawler/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// New repo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// this is handler for home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home function")
	render.RenderTemplatesApproach2(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About function")
	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again"
	render.RenderTemplatesApproach2(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
