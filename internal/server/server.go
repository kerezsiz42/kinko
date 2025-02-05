package server

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/kerezsiz42/kinko/internal/database"
)

type IndexPageData struct {
	Banner  string
	Records []database.Record
}

//go:embed index.template.html
var indexTemplateHTML embed.FS
var tmpl *template.Template

func init() {
	t, err := template.ParseFS(indexTemplateHTML, "index.template.html")
	if err != nil {
		panic(err)
	}

	tmpl = t
}

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	data := IndexPageData{
		Banner: "John Smith's database",
		Records: []database.Record{
			database.NewRecord("Login", "email@email.com", "password", "This is very important"),
			database.NewRecord("Login 2", "email@email.com", "password", "This is very important"),
			database.NewRecord("Login 3", "email@email.com", "password", "This is very important"),
			database.NewRecord("Login 4", "email@email.com", "password", "This is very important"),
		},
	}

	if err := tmpl.Execute(w, data); err != nil {
		slog.Error("Error when executing template", "error", err.Error(), "func", "GetIndexPage")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func PostRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Export(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Import(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", GetIndexPage)
	mux.HandleFunc("POST /records/{id}", PostRecords)
	mux.HandleFunc("GET /records/{id}", PostRecords)
	mux.HandleFunc("GET /records/export", Export)
	mux.HandleFunc("POST /records/import", Import)

	return mux
}
