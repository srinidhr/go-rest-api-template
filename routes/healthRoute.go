package routes

import (
	"log"
	"net/http"
	"time"

	"go-rest-api-template/model"

	"github.com/gin-gonic/gin"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterHealthHandlers(r *gin.RouterGroup, logger log.Logger) {
	res := HealthResource{logger}

	r.GET("/health", res.GetHealthRoute)
	r.GET("/ping", res.GetHealthRoute)
}

type HealthResource struct {
	logger log.Logger
}

func (r HealthResource) GetHealthRoute(ctx *gin.Context) {
	r.logger.Println("Health endpoint hit")

	// Construct health status response
	health := model.Health{
		ServiceName:   "phanes",
		HealthStatus:  "HEALTHY",
		ServiceStatus: "online",
		CurrentTime:   time.Now().UTC(),
	}

	ctx.JSON(http.StatusOK, health)
}
