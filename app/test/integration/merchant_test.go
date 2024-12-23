package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"blackapp/internal/api/router"
	"blackapp/internal/service/dto"
	"blackapp/pkg/config"
	"blackapp/pkg/database"
)

func setupTestRouter(t *testing.T) *gin.Engine {
	// 初始化测试环境
	if err := config.Init(); err != nil {
		t.Fatal(err)
	}

	if err := database.Init(); err != nil {
		t.Fatal(err)
	}

	return router.InitRouter()
}

func TestMerchantAPI_Create(t *testing.T) {
	r := setupTestRouter(t)

	merchant := dto.CreateMerchantDTO{
		Name:          "Test Merchant",
		Address:       "Test Address",
		ContactPerson: "Test Person",
		ContactPhone:  "12345678901",
	}

	body, _ := json.Marshal(merchant)
	req := httptest.NewRequest("POST", "/api/v1/merchants", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), response["code"])
}

func TestMerchantAPI_Login(t *testing.T) {
	r := setupTestRouter(t)

	// 先创建一个商户
	merchant := dto.CreateMerchantDTO{
		Name:          "Test Merchant",
		Address:       "Test Address",
		ContactPerson: "Test Person",
		ContactPhone:  "12345678901",
	}

	body, _ := json.Marshal(merchant)
	req := httptest.NewRequest("POST", "/api/v1/merchants", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 获取商户的API Key和Secret
	// 注意：在实际测试中，你需要从数据库中获取这些信息
	apiKey := "test_key"
	apiSecret := "test_secret"

	// 测试登录
	loginReq := httptest.NewRequest("POST", "/api/v1/merchants/login", nil)
	loginReq.ParseForm()
	loginReq.Form.Add("api_key", apiKey)
	loginReq.Form.Add("api_secret", apiSecret)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, loginReq)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, float64(0), response["code"])
}
