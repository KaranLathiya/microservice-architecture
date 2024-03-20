package portfolio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"portfolio/database"
	"portfolio/model"
	"portfolio/response"
	"sync"
)

func PortfolioList(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	var quit = make(chan error)
	var inputForPortfolios model.InputForPortfolios
	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	err = json.Unmarshal(bodyData, &inputForPortfolios)
	if err != nil {
		response.MessageShow(500, "Internal server error", w)
		return
	}
	expertPortfolioMap := make(map[string]model.PortfoliosOfExpert)
	db := database.GetDB()
	wg.Add(2)
	go portfoliosOfExpert(db, &wg, inputForPortfolios, expertPortfolioMap, w, quit)
	go totalNumberOfPortfoliosOfExpert(db, &wg, inputForPortfolios, expertPortfolioMap, w, quit)
	go func() {
		wg.Wait()
		close(quit)
	}()

	for err := range quit {
		fmt.Println(err)
		return
	}

	result, err := json.MarshalIndent(expertPortfolioMap, "", "  ")
	if err != nil {
		response.MessageShow(400, "Error on marshaling", w)
		return
	}
	w.Write(result)
}
