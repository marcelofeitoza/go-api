package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/marcelofeitoza/track-co-api/internal/domain"
	"log"
)

type NPSService struct {
	DB *sqlx.DB
}

func (n *NPSService) GetNps(widgetID string, npsID string) (domain.NPS, error) {
	nps := domain.NPS{}

	log.Println("Fetching NPS for widget", widgetID, "and NPS", npsID)

	query := `
		SELECT
			id,
			widget_id,
			answer,
			rating,	
			created_at,
			updated_at
		FROM
			nps
		WHERE	
			widget_id = $1
		AND
			id = $2
	`

	rows, err := n.DB.Queryx(query, widgetID, npsID)
	if err != nil {
		log.Println("Error fetching NPS:", err)
		return nps, err
	}

	for rows.Next() {
		err = rows.StructScan(&nps)
		if err != nil {
			log.Println("Error scanning NPS:", err)
			return nps, err
		}
	}

	log.Println("Fetched NPS")
	return nps, nil
}

func (n *NPSService) GetAllNps(widgetID string) ([]domain.NPS, error) {
	nps := []domain.NPS{}

	log.Println("Fetching all NPS for widget", widgetID)

	query := `
		SELECT
			id,
			widget_id,
			answer,
			rating,	
			created_at,
			updated_at
		FROM
			nps
		WHERE	
			widget_id = $1
	`

	rows, err := n.DB.Queryx(query, widgetID)
	if err != nil {
		log.Println("Error fetching NPS:", err)
		return nps, err
	}

	for rows.Next() {
		var n domain.NPS
		err = rows.StructScan(&n)
		if err != nil {
			log.Println("Error scanning NPS:", err)
			return nps, err
		}
		nps = append(nps, n)
	}

	log.Println("Fetched all NPS")
	return nps, nil
}

func (n *NPSService) CreateNps(widgetID string, npsCreateRequest domain.NPSCreateRequest) (domain.NPS, error) {
	nps := domain.NPS{}

	log.Println("Creating NPS for widget", widgetID)

	query := `
		INSERT INTO nps (widget_id, answer, rating)
		VALUES ($1, $2, $3)
		RETURNING id, widget_id, answer, rating, created_at, updated_at
	`

	err := n.DB.QueryRowx(query, widgetID, npsCreateRequest.Answer, npsCreateRequest.Rating).StructScan(&nps)
	if err != nil {
		log.Println("Error creating NPS:", err)
		return nps, err
	}

	log.Println("Created NPS")
	return nps, nil
}
