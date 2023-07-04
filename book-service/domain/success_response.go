package domain

type SuccessfulMessage struct {
	Message string `json:"message"`
}
type BookData struct {
	ID             uint   `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	ChaptersAmount uint   `json:"chapters_amount"`
}
type BookDataSerializator struct {
	Message string   `json:"message"`
	Data    BookData `json:"data"`
}

type BookDataArraySerializator struct {
	Message string     `json:"message"`
	Data    []BookData `json:"data"`
}

type BookProgressSerializator struct {
	Message           string `json:"message"`
	CompletedChapters uint   `json:"completed_chapters"`
	IsBookCompleted   bool   `json:"is_book_completed"`
}
