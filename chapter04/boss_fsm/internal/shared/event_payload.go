package shared

type EventPayload interface{}

type LoginPayload struct {
	UserID  string `json:"userId"`
	IsAdmin bool   `json:"isAdmin"`
}

type LogoutPayload struct {
	UserID string `json:"userId"`
}

type NewMessagePayload struct {
	AuthorID string   `json:"authorId"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}

type NewPostPayload struct {
	PostID   string   `json:"id"`
	AuthorID string   `json:"authorId"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Tags     []string `json:"tags"`
}
