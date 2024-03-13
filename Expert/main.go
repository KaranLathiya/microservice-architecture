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
	r.Route("/user", func(r chi.Router) {
		r.Get("/expert", expertList)
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
	includePortfolios := r.FormValue("includePortfolios")
	includeNumberOfPortfolios := r.FormValue("includeNumberOfPortfolios")
	pageInteger, err := strconv.Atoi(page)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return
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
	includePortfoliosBool, err := strconv.ParseBool(includePortfolios)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return
	}
	if includePortfoliosBool {
		err := includePortfoliosOfExpert(includeNumberOfPortfolios , expertIDs , w , expertsWithPortfolios)
		if err != nil{
			return
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

func includePortfoliosOfExpert(includeNumberOfPortfolios string, expertIDs []string, w http.ResponseWriter, expertsWithPortfolios []model.ExpertWithPortfolio) error {

	includeNumberOfPortfoliosInteger, err := strconv.Atoi(includeNumberOfPortfolios)
	if err != nil {
		response.MessageShow(400, http.StatusText(400), w)
		return err
	}
	if includeNumberOfPortfoliosInteger == 0 {
		if !(includeNumberOfPortfolios == "0"){
		includeNumberOfPortfoliosInteger = 3 
		}
	}
	token, err := createJWT()
	if err != nil {
		fmt.Println(err)
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	fmt.Println("call second service")
	url := "http://localhost:8082/expert/portfolio"
	var inputForPortfolios model.InputForPortfolios
	inputForPortfolios.ExpertIDs = expertIDs
	inputForPortfolios.IncludeNumberOfPortfolios = includeNumberOfPortfoliosInteger
	inputForPortfoliosByte, _ := json.MarshalIndent(inputForPortfolios, "", " ")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(inputForPortfoliosByte))
	if err != nil {
		fmt.Print(err.Error())
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	req.Header.Add("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	var mapOfExpertsWithPortfolios map[string]model.PortfoliosOfExpert
	err = json.Unmarshal(body, &mapOfExpertsWithPortfolios)
	if err != nil {
		fmt.Println(err)
		response.MessageShow(500, http.StatusText(500), w)
		return err
	}
	for i, expertID := range expertIDs {
		expertsWithPortfolios[i].TotalPortfolios = mapOfExpertsWithPortfolios[expertID].TotalPortfolios
		expertsWithPortfolios[i].Portfolios = mapOfExpertsWithPortfolios[expertID].Portfolios
		fmt.Println(expertsWithPortfolios[i])
	}
	return nil
}
