package domain

type SnippetData struct {
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func FetchByIdSuccess(s *Snippet) SnippetData {
	return SnippetData{
		Message: "Snippet is successfully found! =)",
		Data: map[string]any{
			"id":    s.ID,
			"title": s.Title,
			"text":  s.Text,
		},
	}
}
