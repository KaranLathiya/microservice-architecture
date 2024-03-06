package main

import (
	"bytes"
	"encoding/json"
	"expert/dal"
	"expert/model"
	"expert/response"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	jwtKey = []byte(os.Getenv("JWTKEY"))
	db, _ := dal.Connect()
	defer db.Close()
	fmt.Println("Server started")
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/user/expert", expertList)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("wrong method"))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	http.Handle("/", r)
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
	if !(includePortfolios == "false") {
		includeNumberOfPortfoliosInteger, err := strconv.Atoi(includeNumberOfPortfolios)
		if err != nil {
			includeNumberOfPortfoliosInteger = 3
		}
		token, err := createJWT()
		if err != nil {
			fmt.Println(err)
			response.MessageShow(500, "Internal server error", w)
			return
		}
		fmt.Println("call second service")
		url := "http://localhost:8082/user/portfolio"
		var inputForPortfolios model.InputForPortfolios
		inputForPortfolios.ExpertIDs = expertIDs
		inputForPortfolios.IncludeNumberOfPortfolios = includeNumberOfPortfoliosInteger
		inputForPortfoliosByte, _ := json.MarshalIndent(inputForPortfolios, "", " ")
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(inputForPortfoliosByte))
		if err != nil {
			fmt.Print(err.Error())
			response.MessageShow(500, "Internal server error", w)
			return
		}
		req.Header.Add("Authorization", token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Print(err.Error())
			response.MessageShow(500, "Internal server error", w)
			return
		}

		defer res.Body.Close()
		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Print(err.Error())
			response.MessageShow(500, "Internal server error", w)
			return
		}
		var mapOfExpertsWithPortfolios map[string]model.PortfoliosOfExpert
		err = json.Unmarshal(body, &mapOfExpertsWithPortfolios)
		if err != nil {
			fmt.Println(err)
			response.MessageShow(500, "Internal server error", w)
			return
		}
		for i, expertID := range expertIDs {
			expertsWithPortfolios[i].TotalPortfolios = mapOfExpertsWithPortfolios[expertID].TotalPortfolios
			expertsWithPortfolios[i].Portfolios = mapOfExpertsWithPortfolios[expertID].Portfolios
			fmt.Println(expertsWithPortfolios[i])
		}
	}
	result, _ := json.MarshalIndent(expertsWithPortfolios, "", "  ")
	w.Write(result)
}

func createJWT() (string, error) {
	expirationTime := time.Now().Add(time.Minute * 5)

	claims := model.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Audience:  "Portfolio",
			Subject:   "Portfolios of expert",
		},
	}
	// jwt.EncodeSegment(jwtKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
