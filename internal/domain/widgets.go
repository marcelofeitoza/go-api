package domain

type Widget struct {
	ID        int    `db:"id" json:"id"`
	Title     string `db:"title" json:"title"`
	Link      string `db:"link" json:"link"`
	Question  string `db:"question" json:"question"`
	Color     string `db:"color" json:"color"`
	NPS       []NPS  `db:"nps" json:"nps"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type WidgetCreateRequest struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Question string `json:"question"`
}

type WidgetUpdateRequest struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Question string `json:"question"`
	Color    string `json:"color"`
}
