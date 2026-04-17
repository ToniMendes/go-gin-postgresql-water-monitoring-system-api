// Package entities defines the domain entities for the water monitoring system.
package entities

import (
	"fmt"
	"strings"
)

type Owner struct {
	ID        int64
	OwnerName string
	Email     string
	Phone     string
}

func NewOwner(ownerName, email, phone string) (*Owner, error) {
	if strings.TrimSpace(ownerName) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(phone) == "" {
		return &Owner{}, fmt.Errorf("nenhum dos campos devem estar vazios ou conter espaços em branco antes ou no final")
	}

	if len(phone) != 11 {
		return &Owner{}, fmt.Errorf("o número de telefone deve conter 11 dígitos")
	}

	if !strings.Contains(email, "@") && !strings.Contains(email, ".com") {
		return &Owner{}, fmt.Errorf("o Email precisa ser válido")
	}

	return &Owner{
		OwnerName: ownerName,
		Email:     email,
		Phone:     phone,
	}, nil
}
