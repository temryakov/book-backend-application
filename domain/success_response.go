package domain

type SuccessfulMessage struct {
	Message string `json:"message"`
}
type SnippetData struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
type FetchSnippetResponse struct {
	Message string      `json:"message"`
	Data    SnippetData `json:"data"`
}
