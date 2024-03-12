package service

import (
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/marcelofeitoza/track-co-api/internal/domain"
	"log"
)

type WidgetService struct {
	DB *sqlx.DB
}

func (w *WidgetService) GetWidgets() ([]domain.Widget, error) {
	widgets := []domain.Widget{}

	log.Println("Fetching widgets and NPS")

	query := `
		SELECT 
    		w.id, 
		    w.title, 
		    w.link, 
		    w.question, 
    		w.color, 
    		w.created_at, 
   			w.updated_at,
    		COALESCE(json_agg(
        		json_build_object(
            		'id', n.id,
            		'widget_id', n.widget_id,
            		'answer', n.answer,
            		'rating', n.rating,
            		'created_at', n.created_at,
            		'updated_at', n.updated_at
        		) ORDER BY n.created_at) FILTER (WHERE n.id IS NOT NULL), '[]') AS nps
		FROM 
    		widgets w
		LEFT JOIN 
    		nps n ON w.id = n.widget_id
		GROUP BY 
    		w.id
		ORDER BY 
    		w.created_at DESC;
	`

	rows, err := w.DB.Queryx(query)
	if err != nil {
		log.Println("Error fetching widgets:", err)
		return nil, err
	}

	for rows.Next() {
		var widget domain.Widget
		var npsJson []byte

		log.Println("Fetched widgets")

		log.Println("Scanning widget")
		err = rows.Scan(&widget.ID, &widget.Title, &widget.Link, &widget.Question, &widget.Color, &widget.CreatedAt, &widget.UpdatedAt, &npsJson)
		if err != nil {
			log.Println("Error scanning widget:", err)
			return nil, err
		}

		log.Println("Scanned widget")
		err = json.Unmarshal(npsJson, &widget.NPS)
		if err != nil {
			log.Println("Error unmarshalling NPS:", err)
			return nil, err
		}

		log.Println("Unmarshalled NPS")
		widgets = append(widgets, widget)
	}

	return widgets, nil
}

func (w *WidgetService) GetWidget(widgetID string) (domain.Widget, error) {
	widget := domain.Widget{NPS: []domain.NPS{}}

	log.Println("Fetching widget and NPS for widget ID:", widgetID)

	query := `
		SELECT 
    		w.id, 
		    w.title, 
		    w.link, 
		    w.question, 
    		w.color, 
    		w.created_at, 
   			w.updated_at,
    		COALESCE(json_agg(
        		json_build_object(
            		'id', n.id,
            		'widget_id', n.widget_id,
            		'answer', n.answer,
            		'rating', n.rating,
            		'created_at', n.created_at,
            		'updated_at', n.updated_at
        		) ORDER BY n.created_at) FILTER (WHERE n.id IS NOT NULL), '[]') AS nps
		FROM 
    		widgets w
		LEFT JOIN 
    		nps n ON w.id = n.widget_id
		WHERE 
			w.id = $1
		GROUP BY 
    		w.id
		ORDER BY 
    		w.created_at DESC;
	`

	rows, err := w.DB.Queryx(query, widgetID)
	if err != nil {
		log.Println("Error fetching widget:", err)
		return widget, err
	}

	for rows.Next() {
		var npsJson []byte

		log.Println("Fetched widget")

		log.Println("Scanning widget")
		err = rows.Scan(&widget.ID, &widget.Title, &widget.Link, &widget.Question, &widget.Color, &widget.CreatedAt, &widget.UpdatedAt, &npsJson)
		if err != nil {
			log.Println("Error scanning widget:", err)
			return widget, err
		}

		log.Println("Scanned widget")
		err = json.Unmarshal(npsJson, &widget.NPS)
		if err != nil {
			log.Println("Error unmarshalling NPS:", err)
			return widget, err
		}

		log.Println("Unmarshalled NPS")
	}

	return widget, nil
}

func (w *WidgetService) CreateWidget(widgetReq domain.WidgetCreateRequest) (domain.Widget, error) {
	newWidget := domain.Widget{NPS: []domain.NPS{}}

	if widgetReq.Title == "" || widgetReq.Link == "" || widgetReq.Question == "" {
		return newWidget, errors.New("title, link, and question are required")
	}

	err := w.DB.QueryRowx(`
        INSERT INTO widgets (title, link, question)
        VALUES ($1, $2, $3) RETURNING *;
    `, widgetReq.Title, widgetReq.Link, widgetReq.Question).StructScan(&newWidget)
	if err != nil {
		return newWidget, err
	}

	return newWidget, nil
}

func (w *WidgetService) UpdateWidget(widgetID string, widgetReq domain.WidgetUpdateRequest) (domain.Widget, error) {
	updatedWidget := domain.Widget{NPS: []domain.NPS{}}

	if widgetReq.Title == "" && widgetReq.Link == "" && widgetReq.Question == "" && widgetReq.Color == "" {
		return updatedWidget, errors.New("at least one field is required")
	}

	var count int
	_ = w.DB.Get(&count, "SELECT COUNT(*) FROM widgets WHERE id = $1", widgetID)
	if count == 0 {
		return updatedWidget, errors.New("widget not found")
	}

	err := w.DB.QueryRowx(`
		UPDATE widgets
		SET title = COALESCE($1, title), link = COALESCE($2, link), question = COALESCE($3, question), color = COALESCE($4, color)
		WHERE id = $5 RETURNING *;
	`, widgetReq.Title, widgetReq.Link, widgetReq.Question, widgetReq.Color, widgetID).StructScan(&updatedWidget)
	if err != nil {
		return updatedWidget, err
	}

	return updatedWidget, nil
}

func (w *WidgetService) DeleteWidget(widgetID string) (domain.Widget, error) {
	deletedWidget, nps := domain.Widget{NPS: []domain.NPS{}}, []domain.NPS{}

	err := w.DB.Select(&nps, "DELETE FROM nps WHERE widget_id = $1 RETURNING *", widgetID)
	if err != nil {
		log.Println("Error deleting NPS:", err)
		return deletedWidget, err
	}
	log.Println("Deleted NPS:", nps)

	err = w.DB.QueryRowx("DELETE FROM widgets WHERE id = $1 RETURNING *", widgetID).StructScan(&deletedWidget)
	if err != nil {
		log.Println("Error deleting widget:", err)
		return deletedWidget, err
	}
	log.Println("Deleted widget:", deletedWidget)

	deletedWidget.NPS = nps

	log.Println("Deleted widget and NPS:", deletedWidget)
	return deletedWidget, nil
}
