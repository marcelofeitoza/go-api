package routes

import (
	"github.com/marcelofeitoza/track-co-api/internal/api/http/nps"
	"github.com/marcelofeitoza/track-co-api/internal/api/http/widgets"
	"github.com/marcelofeitoza/track-co-api/internal/service"
	"net/http"
)

func NewRouter(app *service.App) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /widgets", widgets.GetWidgetsHandler(app))
	mux.HandleFunc("GET /widgets/{id}", widgets.GetWidgetHandler(app))
	mux.HandleFunc("POST /widgets", widgets.CreateWidgetHandler(app))
	mux.HandleFunc("PUT /widgets/{id}", widgets.UpdateWidgetHandler(app))
	mux.HandleFunc("DELETE /widgets/{id}", widgets.DeleteWidgetHandler(app))

	mux.HandleFunc("GET /widgets/{widgetID}/nps", nps.GetAllNPSHandler(app))
	mux.HandleFunc("GET /widgets/{widgetID}/nps/{npsID}", nps.GetNPSHandler(app))
	mux.HandleFunc("POST /widgets/{widgetID}/nps", nps.CreateNPSHandler(app))

	return mux
}
