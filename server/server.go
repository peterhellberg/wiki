package server

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/peterhellberg/wiki/db"
	"github.com/russross/blackfriday"
)

// Server is the wiki server
type Server struct {
	logger *log.Logger
	db     *db.DB
}

// New creates a new wiki server
func New(logger *log.Logger, db *db.DB) *Server {
	return &Server{logger: logger, db: db}
}

func (s *Server) redirect(w http.ResponseWriter, r *http.Request) {
	if path := r.URL.Path; len(r.URL.Path) > 1 {
		target := strings.TrimSuffix(path, "/")

		if target == "/home" {
			target = "/"
		}

		http.Redirect(w, r, target, 302)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "wiki")

	s.logger.Println(r.Method, r.URL.String())

	switch {
	case r.Method == http.MethodPost:
		s.save(w, r)
	case r.URL.Path == "/favicon.ico":
		s.favicon(w, r)
	case r.URL.Path == "/home":
		s.redirect(w, r)
	case strings.HasSuffix(r.URL.Path, "/edit"):
		s.edit(w, r)
	case strings.HasSuffix(r.URL.Path, "/") && len(r.URL.Path) > 1:
		s.redirect(w, r)
	default:
		s.show(w, r)
	}
}

func (s *Server) getPageName(r *http.Request) []byte {
	name := strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/edit"), "/")

	if name == "" {
		return []byte("home")
	}

	return []byte(name)
}

type data map[string]interface{}

// bytesAsHTML returns the template bytes as HTML
func bytesAsHTML(b []byte) template.HTML {
	return template.HTML(string(b))
}

// parsedMarkdown returns provided bytes parsed as Markdown
func parsedMarkdown(b []byte) []byte {
	return blackfriday.MarkdownCommon(b)
}
