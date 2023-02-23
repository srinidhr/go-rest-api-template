package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-rest-api-template/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockHealthResoure = HealthResource{*log.Default()}
var mockHealthRouter = setupMockHealthRouter()

func TestGetHealthRoute(t *testing.T) {
	postRequest, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()
	mockHealthRouter.ServeHTTP(w, postRequest)
	assert.Equal(t, http.StatusOK, w.Code)

	log.Println(w.Body)

	var healthResponse = &model.Health{}
	json.Unmarshal(w.Body.Bytes(), healthResponse)

	assert.Equal(t, healthResponse.ServiceName, "phanes")
	assert.Equal(t, healthResponse.ServiceStatus, "online")
	assert.Equal(t, healthResponse.HealthStatus, "HEALTHY")
	assert.IsType(t, healthResponse.CurrentTime, time.Time{})
}

func TestRegisterHealthHandlers(t *testing.T) {
	RegisterHealthHandlers(gin.Default().Group(""), *log.Default())
}

func setupMockHealthRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	r.GET("/health", mockHealthResoure.GetHealthRoute)
	r.GET("/ping", mockHealthResoure.GetHealthRoute)

	return r
}
