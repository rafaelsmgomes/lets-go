package main

import "snippetbox.rafaelsmgomes/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
