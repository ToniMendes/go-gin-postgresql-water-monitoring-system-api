package entities

type Address struct {
	CEP          string
	PublicPlace  string
	Neighborhood string
	State        string
	City         string
	Region       string
}

func NewAddress(cep, publicPlace, neighborhood, state, city, region string) *Address {
	return &Address{
		CEP:          cep,
		PublicPlace:  publicPlace,
		Neighborhood: neighborhood,
		State:        state,
		City:         city,
	}
}