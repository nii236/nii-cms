package handlers

import "html/template"

func loadTemplate(path string) (*template.Template, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return t, err
}
