package main

import (
	"encoding/json"
	"expert/dal"
	"expert/model"
	"expert/response"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_ke")

func main() {
	db, _ := dal.Connect()
	defer db.Close()
	fmt.Println("Server started")
	http.HandleFunc("/expert", expertList)
	http.ListenAndServe(":8081", nil)
}

func expertList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page := r.FormValue("page")
	page = strings.TrimSpace(page)
	includePortfolios := r.FormValue("includePortfolios")
	includePortfolios = strings.TrimSpace(includePortfolios)
	includeNumberOfPortfolios := r.FormValue("includeNumberOfPortfolios")
	includeNumberOfPortfolios = strings.TrimSpace(includeNumberOfPortfolios)
	pageInteger, err := strconv.Atoi(page)
	if err != nil {
		pageInteger = 0
	}
	db := dal.GetDB()
	rows, err := db.Query("SELECT id,name FROM public.expert order by name LIMIT 10 OFFSET $1", pageInteger*10)
	if err != nil {
		databaseErrorMessage, databaseErrorCode := response.DatabaseErrorShow(err)
		response.MessageShow(databaseErrorCode, databaseErrorMessage, w)
		return
	}
	var expertsWithPortfolios []model.ExpertWithPortfolio
	expertIDs := []string{}
	fmt.Println(expertIDs)
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
	fmt.Println(expertsWithPortfolios)
	if !(includePortfolios == "false") {
		includeNumberOfPortfoliosInteger, err := strconv.Atoi(includeNumberOfPortfolios)
		if err != nil {
			includeNumberOfPortfoliosInteger = 3
		}
		token, err := createJWT(expertIDs, includeNumberOfPortfoliosInteger)
		if err != nil {
			fmt.Println(err)
			response.MessageShow(500, "Internal server error", w)
			return
		}
		fmt.Println("call second service")
		url := "http://localhost:8082/user/portfolio"
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Print(err.Error())
		}
		req.Header.Add("Authorization", token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Print(err.Error())
		}

		defer res.Body.Close()
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Print(err.Error())
		}
		var mapOfExpertsWithPortfolios map[string]model.PortfoliosOfExpert
		err = json.Unmarshal(body, &mapOfExpertsWithPortfolios)
		if err != nil {
			fmt.Println(err)
			response.MessageShow(500, "Internal server error", w)
			return
		}
		fmt.Println(mapOfExpertsWithPortfolios)
		for i, expertID := range expertIDs {
			expertsWithPortfolios[i].TotalPortfolios = mapOfExpertsWithPortfolios[expertID].TotalPortfolios
			expertsWithPortfolios[i].Portfolios = mapOfExpertsWithPortfolios[expertID].Portfolios
			fmt.Println(expertsWithPortfolios[i])
		}
	}
	result, _ := json.MarshalIndent(expertsWithPortfolios, "", "  ")
	w.Write(result)
}

func createJWT(expertIDs []string, includeNumberOfPortfolios int) (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := model.Claims{
		ExpertID:                  expertIDs,
		IncludeNumberOfPortfolios: includeNumberOfPortfolios,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Audience:  "Portfolio",
			Subject:   "Portfolios of expert",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 	i := 1
// 	for i < 10000 {
// 		name := "user" + strconv.Itoa(i)
// 		_ = db.QueryRow("INSERT INTO public.expert (name) VALUES ($1);", name)
// 		// if err != nil {
// 		// 	fmt.Println(err)
// 		// 	break
// 		// }
// 		i+=1
// 	}
// }

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func RandStringRunes(n int) string {
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }
