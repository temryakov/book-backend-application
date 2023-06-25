package domain

type SuccessfulMessage struct {
	Message string `json:"message"`
}
type BookData struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
type BookDataSerializator struct {
	Message string   `json:"message"`
	Data    BookData `json:"data"`
}

type BookDataArraySerializator struct {
	Message string     `json:"message"`
	Data    []BookData `json:"data"`
}
