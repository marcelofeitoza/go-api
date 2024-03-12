package nps

import (
	"encoding/json"
	"github.com/marcelofeitoza/track-co-api/internal/domain"
	"github.com/marcelofeitoza/track-co-api/internal/service"
	"log"
	"net/http"
)

func GetNPSHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("widgetID")
		npsId := r.PathValue("npsId")
		npsService := service.NPSService{DB: app.DB}

		log.Println("Fetching NPS", npsId, "for widget", id)

		nps, err := npsService.GetNps(id, npsId)
		if err != nil {
			log.Println("Error fetching NPS:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Fetched NPS")

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(nps); err != nil {
			log.Println("Error encoding NPS into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func GetAllNPSHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("widgetID")
		npsService := service.NPSService{DB: app.DB}

		log.Println("Fetching all NPS for widget", id)

		nps, err := npsService.GetAllNps(id)
		if err != nil {
			log.Println("Error fetching NPS:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Fetched NPS")

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(nps); err != nil {
			log.Println("Error encoding NPS into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func CreateNPSHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("widgetID")
		var npsReq domain.NPSCreateRequest
		npsService := service.NPSService{DB: app.DB}

		err := json.NewDecoder(r.Body).Decode(&npsReq)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		newNPS, err := npsService.CreateNps(id, npsReq)
		if err != nil {
			log.Println("Error creating NPS:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Created NPS", newNPS.ID)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(newNPS); err != nil {
			log.Println("Error encoding NPS into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
