package config

import (
	scs "github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
