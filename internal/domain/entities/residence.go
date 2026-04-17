package entities

type Residence struct {
	ID           int64   `db:"id"`
	OwnerName    string  `db:"owners_name"`
	Email        string  `db:"email"`
	Phone        string  `db:"phone"`
	CEP          string  `db:"cep"`
	PublicPlace  string  `db:"publicplace"`
	Neighborhood string  `db:"neighborhood"`
	Uf           string  `db:"uf"`
	City         string  `db:"city"`
	Region       string  `db:"region"`
	Invoice      float64 `db:"invoice"`
}
