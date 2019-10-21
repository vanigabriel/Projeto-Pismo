package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evalphobia/go-timber/timber"

	"github.com/stretchr/testify/assert"

	"github.com/vanigabriel/Projeto-Pismo/entity"
	"github.com/vanigabriel/Projeto-Pismo/usecases"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateAccount(t *testing.T) {
	// Prepara para testar em memória
	s := usecases.NewService(usecases.InitStatements())

	conf := timber.Config{
		APIKey:         "",
		SourceID:       "",
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelInfo,
		Sync:           true,
		Debug:          true,
	}

	log, _ := timber.New(conf)

	router := SetupRouter(s, log)

	// Insert Correto
	acc := &Account{
		CreditLimit: Amount{
			Value: 2000.00,
		},
		WithdrawlLimit: Amount{
			Value: 2000.00,
		},
	}

	b, _ := json.Marshal(acc)
	send := bytes.NewBuffer(b)

	// Performa Post
	w := performRequest(router, "POST", "/v1/accounts", send)

	// Valida se retornou 201
	assert.Equal(t, http.StatusCreated, w.Code)

	// Converte resposta
	var response entity.Account
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Validações
	assert.Nil(t, err)
	assert.Equal(t, response.CreditLimit, acc.CreditLimit.Value)

}

func TestAlterAccount(t *testing.T) {
	// Prepara para testar em memória
	s := usecases.NewService(usecases.InitStatements())

	conf := timber.Config{
		APIKey:         "",
		SourceID:       "",
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelInfo,
		Sync:           true,
		Debug:          true,
	}

	log, _ := timber.New(conf)

	router := SetupRouter(s, log)

	// Nova conta
	acc := &Account{
		CreditLimit: Amount{
			Value: 2000.00,
		},
		WithdrawlLimit: Amount{
			Value: 2000.00,
		},
	}

	b, _ := json.Marshal(acc)
	send := bytes.NewBuffer(b)

	// Cria Conta
	w := performRequest(router, "POST", "/v1/accounts", send)

	patchAccount := &Account{
		CreditLimit: Amount{
			Value: 200.00,
		},
		WithdrawlLimit: Amount{
			Value: 200.00,
		},
	}

	b, _ = json.Marshal(patchAccount)
	send = bytes.NewBuffer(b)

	// Atualiza
	w = performRequest(router, "PATCH", "/v1/accounts/1", send)

	// Valida se retornou 201
	assert.Equal(t, http.StatusCreated, w.Code)

	// Converte resposta
	var response entity.Account
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// Validações
	assert.Nil(t, err)
	assert.Equal(t, response.CreditLimit, 2200.00)

	// Tentativa de atualizar uma conta inexistente
	w = performRequest(router, "PATCH", "/v1/accounts/8", send)

	// Valida se retornou 400
	assert.Equal(t, 400, w.Code)

	// Converte resposta para um MAP
	var response2 map[string]string
	err = json.Unmarshal([]byte(w.Body.String()), &response2)

	// Recupera a tag error e verifica se existe
	_, exists := response2["error"]

	// Validações
	assert.Nil(t, err)
	assert.True(t, exists)

}
