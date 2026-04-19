package entities_test

import (
	"go-gin-postgresql-water-monitoring-system-api/internal/domain/entities"
	"testing"
)

func TestNewOwner(t *testing.T) {
	type args struct {
		ownerName string
		email     string
		phone     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Sucesso - dados válidos",
			args: args{
				ownerName: "João Silva",
				email:     "joao@exemplo.com",
				phone:     "11987654321",
			},
			wantErr: false,
		},
		{
			name: "Erro - campos vazios",
			args: args{
				ownerName: " ",
				email:     "joao@exemplo.com",
				phone:     "11987654321",
			},
			wantErr: true,
		},
		{
			name: "Erro - telefone com tamanho inválido",
			args: args{
				ownerName: "João Silva",
				email:     "joao@exemplo.com",
				phone:     "123",
			},
			wantErr: true,
		},
		{
			name: "Erro - email inválido",
			args: args{
				ownerName: "João Silva",
				email:     "email-sem-arroba",
				phone:     "11987654321",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := entities.NewOwner(tt.args.ownerName, tt.args.email, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOwner() erro = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
