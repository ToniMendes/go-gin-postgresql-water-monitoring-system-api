// Package entities defines the domain entities for the water monitoring system.
package entities

type Owner struct {
	ID        int64
	OwnerName string
	Email     string
	Phone     string
}

func NewOwner(ownerName, email, phone string) *Owner {
	return &Owner{
		OwnerName: ownerName,
		Email:     email,
		Phone:     phone,
	}
}
