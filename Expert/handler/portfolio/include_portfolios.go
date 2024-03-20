package portfolio

import (
	"expert/Auth/jwt"
	"expert/model"
	"expert/response"
	"fmt"
	"net/http"
	"strconv"
)

func IncludePortfoliosOfExpert(includeNumberOfPortfolios string, expertIDs []string, w http.ResponseWriter, expertsWithPortfolios []model.ExpertWithPortfolio) error {
	includeNumberOfPortfoliosInteger, err := strconv.Atoi(includeNumberOfPortfolios)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return err
	}
	if includeNumberOfPortfoliosInteger == 0 {
		if !(includeNumberOfPortfolios == "0") {
			includeNumberOfPortfoliosInteger = 3
		}
	}
	jwtToken, err := jwt.CreateJWT()
	if err != nil {
		fmt.Println(err)
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	var inputForPortfolios model.InputForPortfolios
	inputForPortfolios.ExpertIDs = expertIDs
	inputForPortfolios.IncludeNumberOfPortfolios = includeNumberOfPortfoliosInteger
	err = portfolioService(jwtToken, inputForPortfolios, expertsWithPortfolios, w)
	return err
}
