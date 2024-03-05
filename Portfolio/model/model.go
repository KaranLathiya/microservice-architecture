package model

import "github.com/dgrijalva/jwt-go"

type Message struct {
	Code    int    `json:"code"  validate:"required"`
	Message string `json:"message"  validate:"required"`
}

type Claims struct {
	ExpertID                  []string `json:"expertID"`
	IncludeNumberOfPortfolios int      `json:"includeNumberOfPortfolios"`
	jwt.StandardClaims
}

type Portfolio struct {
	ID        string `json:"ID"  validate:"required"`
	Name      string `json:"name"  validate:"required"`
	CreatedAt string `json:"ceatedAt"  validate:"required"`
	Image     string `json:"image"  validate:"required"`
}

type PortfoliosOfExpert struct {
	TotalPortfolios int16       `json:"totalPortfolios"  validate:"required"`
	Portfolios      []Portfolio `json:"portfolios,omitempty" `
}
