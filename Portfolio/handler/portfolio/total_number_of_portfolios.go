package portfolio

import (
	"database/sql"
	"fmt"
	"net/http"
	"portfolio/model"
	"portfolio/response"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx"
)

func totalNumberOfPortfoliosOfExpert(db *sql.DB, wg *sync.WaitGroup, inputForPortfolios model.InputForPortfolios, expertPortfolioMap map[string]model.PortfoliosOfExpert, w http.ResponseWriter, quit chan<- error) {
	defer wg.Done()
	query := fmt.Sprintf(`SELECT created_by, count(id) as total_portfolios FROM public.portfolio where created_by in ('%v')  group by created_by ;`, strings.Join(inputForPortfolios.ExpertIDs, "' , '"))
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := db.Query(query)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		quit <- err
		return
	}
	for rows.Next() {
		var totalPortfolios int
		var expertID string
		err = rows.Scan(&expertID, &totalPortfolios)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			quit <- err
			return
		}
		if expertPortfoliosMapValue, ok := expertPortfolioMap[expertID]; ok {
			expertPortfoliosMapValue.TotalPortfolios = int16(totalPortfolios)
			expertPortfolioMap[expertID] = expertPortfoliosMapValue
		} else {
			var portfoliosOfExpert model.PortfoliosOfExpert
			portfoliosOfExpert.TotalPortfolios = int16(totalPortfolios)
			expertPortfolioMap[expertID] = portfoliosOfExpert
		}
	}
}
