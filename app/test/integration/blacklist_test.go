package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"blackapp/internal/service/dto"
)

func TestBlacklistAPI_Create(t *testing.T) {
	r := setupTestRouter(t)

	blacklist := dto.CreateBlacklistDTO{
		Name:       "Test User",
		Phone:      "12345678901",
		IDCard:     "123456789012345678",
		MerchantID: 1,
	}

	body, _ := json.Marshal(blacklist)
	req := httptest.NewRequest("POST", "/api/v1/blacklists", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	// 添加认证token
	req.Header.Set("Authorization", "Bearer test_token")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), response["code"])
}

func TestBlacklistAPI_Check(t *testing.T) {
	r := setupTestRouter(t)

	check := dto.CheckBlacklistDTO{
		Phone:  "12345678901",
		IDCard: "123456789012345678",
		Name:   "Test User",
	}

	body, _ := json.Marshal(check)
	req := httptest.NewRequest("POST", "/api/v1/blacklists/check", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer test_token")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), response["code"])
}
