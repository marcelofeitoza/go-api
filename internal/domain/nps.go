package domain

type NPS struct {
	ID        int    `db:"id" json:"id"`
	WidgetId  int    `db:"widget_id" json:"widget_id"`
	Answer    string `db:"answer" json:"answer"`
	Rating    int    `db:"rating" json:"rating"`
	CreatedAt string `db:"created_at" json:"created_at"`
	UpdatedAt string `db:"updated_at" json:"updated_at"`
}

type NPSCreateRequest struct {
	Answer string `json:"answer"`
	Rating int    `json:"rating"`
}
