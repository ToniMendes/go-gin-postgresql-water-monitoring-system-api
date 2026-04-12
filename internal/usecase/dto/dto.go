// Package dto usecase
package dto

type WaterMonitoringInput struct {
	OwnerName string `json:"owner"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CEP       string `json:"cep"`
}

type WaterMonitoringOutput struct {
	ID           int    `json:"id"`
	OwnerName    string `json:"owner"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	CEP          string `json:"cep"`
	PublicPlace  string `json:"public_place"`
	Neighborhood string `json:"neighborhood"`
	State        string `json:"state"`
	City         string `json:"city"`
	Region       string `json:"region"`
}
