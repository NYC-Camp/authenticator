// Package for the templating system for Authenticator
package libtmpl

import (
	"html/template"
	"net/http"
)

//@TODO: Change the Content string to be a slice of strings, use those strings
//when executing the template so more robust templates can be supported.
type HTMLTemplateConfig struct {
	Template         HTMLTemplate
	TemplateDir      string
	DefaultErrorFunc TemplateErrorFunc
}

type HTMLTemplate struct {
	Base      string
	BaseName  string
	Head      string
	Header    string
	Footer    string
	Content   string
	ErrorFunc TemplateErrorFunc
}

type HTMLTemplateData struct {
	Head    map[string]interface{}
	Header  map[string]interface{}
	Content map[string]interface{}
	Footer  map[string]interface{}
}

func (htc HTMLTemplateConfig) NewHTMLTemplate() HTMLTemplate {
	templateBase := htc.TemplateDir
	htmlTemplate := HTMLTemplate{}
	htmlTemplate.Base = templateBase + "base.html"
	htmlTemplate.BaseName = "base"
	htmlTemplate.Head = templateBase + "head.html"
	htmlTemplate.Header = templateBase + "header.html"
	htmlTemplate.Footer = templateBase + "footer.html"
	htmlTemplate.ErrorFunc = htc.DefaultErrorFunc

	return htmlTemplate
}

// Function called when an error occurs while executing a template
type TemplateErrorFunc func(w http.ResponseWriter, err error)

func (ht HTMLTemplate) Execute(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(ht.Content, ht.Base, ht.Head, ht.Footer, ht.Header)
	if err != nil {
		ht.ErrorFunc(w, err)
		return
	}

	t.ExecuteTemplate(w, ht.BaseName, data)
}
