package app

import (
	"github.com/kosha/freshdesk-connector/pkg/httpclient"
	"github.com/kosha/freshdesk-connector/pkg/models"
	"net/http"
)

// searchTickets godoc
// @Summary Search tickets
// @Description Search tickets by passing various query parameters
// @Tags search
// @Accept  json
// @Produce  json
// @Param query query string true "Enter query parameter e.g., priority:>2 (in quotes)"
// @Success 200 {object} models.SearchResults
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/search [get]
func (a *App) searchTickets(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	query := r.FormValue("query")
	page := r.FormValue("page")

	t := httpclient.SearchTickets(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey(), query, page)
	for _, ticket := range t.Results {
		status := ticket.Status.(float64)
		ticket.Status = models.Status(status).String()

		priority := ticket.Priority.(float64)
		ticket.Priority = models.Priority(priority).String()

		source := ticket.Source.(float64)
		ticket.Source = models.Source(source).String()
	}
	respondWithJSON(w, http.StatusOK, t)
	return

}
