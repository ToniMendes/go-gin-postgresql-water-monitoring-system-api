// Package viacep provides integration with the ViaCEP API for address lookups.
package viacep

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type viacepInputDTO struct {
	PublicPlace  string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Region       string `json:"regiao"`
	Erro         bool   `json:"erro"`
}

func NewQuery(cep string) (*viacepInputDTO, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data viacepInputDTO
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Erro {
		return nil, fmt.Errorf("CEP inválido")
	}

	return &data, nil
}
