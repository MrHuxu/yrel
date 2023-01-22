package templates

import (
	"embed"
	"fmt"
	"html/template"
)

//go:embed *
var templatesFS embed.FS

// GetTemplate ...
func GetTemplate(name string) (*template.Template, error) {
	filename := fmt.Sprintf("%s.tmpl", name)
	tmpl, err := template.New(filename).ParseFS(templatesFS, filename)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
