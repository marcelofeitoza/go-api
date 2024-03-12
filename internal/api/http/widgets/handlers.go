package widgets

import (
	"encoding/json"
	"github.com/marcelofeitoza/track-co-api/internal/domain"
	"github.com/marcelofeitoza/track-co-api/internal/service"
	"log"
	"net/http"
)

func GetWidgetsHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Fetching widgets")
		widgetService := service.WidgetService{DB: app.DB}

		widgets, err := widgetService.GetWidgets()

		if err != nil {
			log.Println("Error fetching widgets:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Fetched widgets")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(widgets); err != nil {
			log.Println("Error encoding widgets into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func GetWidgetHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		widgetID := r.PathValue("id")
		widgetService := service.WidgetService{DB: app.DB}

		log.Println("Fetching widget")
		widget, err := widgetService.GetWidget(widgetID)

		if err != nil {
			log.Println("Error fetching widget:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Fetched widget")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(widget); err != nil {
			log.Println("Error encoding widget into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func CreateWidgetHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var widgetReq domain.WidgetCreateRequest
		widgetService := service.WidgetService{DB: app.DB}

		err := json.NewDecoder(r.Body).Decode(&widgetReq)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		newWidget, err := widgetService.CreateWidget(widgetReq)
		if err != nil {
			log.Println("Error creating widget:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		log.Println("Created widget", newWidget.ID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(newWidget); err != nil {
			log.Println("Error encoding widget into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func UpdateWidgetHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		widgetID := r.PathValue("id")
		var widgetReq domain.WidgetUpdateRequest
		widgetService := service.WidgetService{DB: app.DB}

		err := json.NewDecoder(r.Body).Decode(&widgetReq)
		if err != nil {
			log.Println("Error decoding request body:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		updatedWidget, err := widgetService.UpdateWidget(widgetID, widgetReq)
		if err != nil {
			log.Println("Error updating widget:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(updatedWidget); err != nil {
			log.Println("Error encoding widget into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func DeleteWidgetHandler(app *service.App) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		widgetID := r.PathValue("id")
		widgetService := service.WidgetService{DB: app.DB}

		deletedWidget, err := widgetService.DeleteWidget(widgetID)
		if err != nil {
			log.Println("Error deleting widget:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(deletedWidget); err != nil {
			log.Println("Error encoding widget into JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
