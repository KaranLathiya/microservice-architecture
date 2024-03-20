package portfolio

import (
	"bytes"
	"encoding/json"
	"expert/model"
	"expert/response"
	"io"
	"net/http"
)

func portfolioService(jwtToken string, inputForPortfolios model.InputForPortfolios, expertsWithPortfolios []model.ExpertWithPortfolio, w http.ResponseWriter) error {
	url := "http://localhost:8082/expert/portfolio"
	inputForPortfoliosByte, _ := json.MarshalIndent(inputForPortfolios, "", " ")
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(inputForPortfoliosByte))
	if err != nil {
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	req.Header.Add("Authorization", jwtToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	var mapOfExpertsWithPortfolios map[string]model.PortfoliosOfExpert
	err = json.Unmarshal(body, &mapOfExpertsWithPortfolios)
	if err != nil {
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	for i, expertID := range inputForPortfolios.ExpertIDs {
		expertsWithPortfolios[i].TotalPortfolios = mapOfExpertsWithPortfolios[expertID].TotalPortfolios
		expertsWithPortfolios[i].Portfolios = mapOfExpertsWithPortfolios[expertID].Portfolios
	}
	return nil
}
