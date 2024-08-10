package note

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
