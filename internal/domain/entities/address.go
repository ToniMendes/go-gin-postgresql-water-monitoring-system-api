package entities

type Address struct {
	CEP          string
	PublicPlace  string
	Neighborhood string
	Uf           string
	City         string
	Region       string
}

func NewAddress(cep, publicPlace, neighborhood, state, city, region string) *Address {
	return &Address{
		CEP:          cep,
		PublicPlace:  publicPlace,
		Neighborhood: neighborhood,
		Uf:           state,
		City:         city,
		Region:       region,
	}
}
