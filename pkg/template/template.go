package template

import (
	"io"
	"io/fs"
	"net/http"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TemplateService interface {
	Execute(w io.Writer, name string, data any) error
	Redirect(w http.ResponseWriter, r *http.Request, url string, code int)
}

type templateService struct {
	templates *template.Template
}

func NewTemplateService(fs fs.FS, templatesPattern string) TemplateService {
	t := template.Must(template.New("").Funcs(templateFunctions).ParseFS(fs, templatesPattern))

	return &templateService{
		templates: t,
	}
}

func (t *templateService) Execute(w io.Writer, name string, data any) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *templateService) Redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

var templateFunctions = template.FuncMap{
	"titleCase": cases.Title(language.BritishEnglish).String,
}
