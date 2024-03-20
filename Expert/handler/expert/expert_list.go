package expert

import (
	"encoding/json"
	"expert/database"
	"expert/handler/portfolio"
	"expert/model"
	"expert/response"
	"net/http"
	"strconv"
)

func ExpertList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := r.FormValue("page")
	includePortfoliosString := r.FormValue("includePortfolios")
	includeNumberOfPortfolios := r.FormValue("includeNumberOfPortfolios")
	pageInteger, err := strconv.Atoi(page)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return
	}
	db := database.GetDB()
	rows, err := db.Query("SELECT id,name FROM public.expert order by name LIMIT 10 OFFSET $1", pageInteger*10)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	var expertsWithPortfolios []model.ExpertWithPortfolio
	expertIDs := []string{}
	for rows.Next() {
		var expertWithPortfolio model.ExpertWithPortfolio
		err = rows.Scan(&expertWithPortfolio.ID, &expertWithPortfolio.Name)
		if err != nil {
			response.MessageShow(400, err.Error(), w)
			return
		}
		expertIDs = append(expertIDs, expertWithPortfolio.ID)
		expertsWithPortfolios = append(expertsWithPortfolios, expertWithPortfolio)
	}
	includePortfolios, err := strconv.ParseBool(includePortfoliosString)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return
	}
	if includePortfolios {
		err := portfolio.IncludePortfoliosOfExpert(includeNumberOfPortfolios, expertIDs, w, expertsWithPortfolios)
		if err != nil {
			return
		}
	}
	result, err := json.MarshalIndent(expertsWithPortfolios, "", "  ")
	if err != nil {
		response.MessageShow(400, "Error on marshaling", w)
		return
	}
	w.Write(result)
}
