package server

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"text/template"
)

type IndexPageData struct {
	Message string
}

//go:embed template.html
var templateHTML embed.FS
var tmpl *template.Template

func init() {
	t, err := template.ParseFS(templateHTML, "template.html")
	if err != nil {
		panic(err)
	}

	tmpl = t
}

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	data := IndexPageData{
		Message: "Hello, this is a simple HTML template example in Go!",
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

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", GetIndexPage)
	mux.HandleFunc("POST /records/{id}", PostRecords)
	mux.HandleFunc("POST /export", Export)

	return mux
}
