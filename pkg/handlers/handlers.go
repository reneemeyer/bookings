package handlers

import (
	"github.com/reneemeyer/bookings/pkg/config"
	"github.com/reneemeyer/bookings/pkg/models"
	"github.com/reneemeyer/bookings/pkg/render"
	"log"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform biz logic
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	log.Println(remoteIP, "remote ip address in about")
	stringMap := map[string]string{
		"test":      "Hello, again",
		"remote_ip": remoteIP,
	}
	log.Println(stringMap, "STRING MAP SENT")
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
