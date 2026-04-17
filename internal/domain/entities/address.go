package entities

import (
	"fmt"
	"strings"
)

type Address struct {
	CEP          string
	PublicPlace  string
	Neighborhood string
	Uf           string
	City         string
	Region       string
}

func NewAddress(cep, publicPlace, neighborhood, state, city, region string) (*Address, error) {
	if (strings.TrimSpace(cep) == "" || len(cep) != 8 || strings.Contains(cep, "-")) || strings.TrimSpace(publicPlace) == "" || strings.TrimSpace(neighborhood) == "" || strings.TrimSpace(state) == "" || strings.TrimSpace(city) == "" || strings.TrimSpace(region) == "" {
		return &Address{}, fmt.Errorf("cep invalido, CEP não deve conter traços ou ser menor que 8 caracteres")
	}

	return &Address{
		CEP:          cep,
		PublicPlace:  publicPlace,
		Neighborhood: neighborhood,
		Uf:           state,
		City:         city,
		Region:       region,
	}, nil
}
