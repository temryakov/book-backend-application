package domain

type SuccessResponse struct {
	Message string `json:"message"`
}
type SnippetData struct {
	Message string `json:"message"`
	Data    struct {
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"data"`
}

func FetchByIdSuccess(s *Snippet) SnippetData {
	return SnippetData{
		Message: "Snippet is successfully found! =)",
		Data: struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
			Text  string `json:"text"`
		}{
			ID:    s.ID,
			Title: s.Title,
			Text:  s.Text,
		},
	}
}
