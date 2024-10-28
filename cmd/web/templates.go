package main

import "github.com/kharljhon14/snippetbox/internal/models"

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
