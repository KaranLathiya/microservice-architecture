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

func portfoliosOfExpert(db *sql.DB, wg *sync.WaitGroup, inputForPortfolios model.InputForPortfolios, expertPortfolioMap map[string]model.PortfoliosOfExpert, w http.ResponseWriter, quit chan<- error) {
	defer wg.Done()
	query := fmt.Sprintf(`WITH RankedPortfolios AS (
		SELECT id, name, created_by, image, created_at, ROW_NUMBER() OVER( partition  by created_by order by created_at desc) row_num
		FROM public.portfolio where created_by in ('%v') order by created_at desc )
		select id, name, created_by, image, created_at from RankedPortfolios where row_num <= %d;`, strings.Join(inputForPortfolios.ExpertIDs, "' , '"), inputForPortfolios.IncludeNumberOfPortfolios)
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := db.Query(query)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		quit <- err
		return
	}
	for rows.Next() {
		var portfolio model.Portfolio
		var expertID string
		err = rows.Scan(&portfolio.ID, &portfolio.Name, &expertID, &portfolio.Image, &portfolio.CreatedAt)
		if err != nil {
			databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
			response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
			quit <- err
			return
		}

		if expertPortfoliosMapValue, ok := expertPortfolioMap[expertID]; ok {
			expertPortfoliosMapValue.Portfolios = append(expertPortfoliosMapValue.Portfolios, portfolio)
			expertPortfolioMap[expertID] = expertPortfoliosMapValue
		} else {
			var portfoliosOfExpert model.PortfoliosOfExpert
			portfoliosOfExpert.Portfolios = append(portfoliosOfExpert.Portfolios, portfolio)
			expertPortfolioMap[expertID] = portfoliosOfExpert
		}
	}
}
