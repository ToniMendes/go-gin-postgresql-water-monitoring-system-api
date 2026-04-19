package web_test

import (
	"bytes"
	"encoding/json"
	"go-gin-postgresql-water-monitoring-system-api/internal/usecase/dto"
	"go-gin-postgresql-water-monitoring-system-api/internal/web"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// mockUseCase implementa a interface UseCaseRepository para testes
type mockUseCase struct {
	execCreateFunc func(dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error)
}

func (m *mockUseCase) ExecCreate(i dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error) {
	return m.execCreateFunc(i)
}

func (m *mockUseCase) ExecUpdateOwner(i dto.WaterMonitoringInput, id int64) (dto.WaterMonitoringOutput, error) {
	return dto.WaterMonitoringOutput{}, nil
}

func TestHandler_AddNewResidence(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("deve retornar 201 quando a criação for bem sucedida", func(t *testing.T) {
		// Setup do mock
		mockUC := &mockUseCase{
			execCreateFunc: func(i dto.WaterMonitoringInput) (dto.WaterMonitoringOutput, error) {
				return dto.WaterMonitoringOutput{ID: 1, OwnerName: i.OwnerName}, nil
			},
		}
		h := web.NewHandler(mockUC)

		// Recorder para capturar a resposta
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		// Dados de entrada
		inputData := map[string]string{
			"owner": "Usuario Teste",
			"email": "teste@teste.com",
			"phone": "11999999999",
			"cep":   "01001000",
		}
		body, _ := json.Marshal(inputData)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/residences", bytes.NewBuffer(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		// Execução
		h.AddNewResidence(ctx)

		// Asserts
		if w.Code != http.StatusCreated {
			t.Errorf("esperado status 201, obtido %v", w.Code)
		}

		var resp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &resp)
		if resp["OwnerName"] != "Usuario Teste" {
			t.Errorf("esperado nome 'Usuario Teste', obtido %v", resp["OwnerName"])
		}
	})
}
