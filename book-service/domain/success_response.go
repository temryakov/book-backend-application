package domain

type SuccessfulMessage struct {
	Message string `json:"message"`
}

type BookProgressSerializator struct {
	Message           string `json:"message"`
	CompletedChapters uint   `json:"completed_chapters"`
	IsBookCompleted   bool   `json:"is_book_completed"`
}
