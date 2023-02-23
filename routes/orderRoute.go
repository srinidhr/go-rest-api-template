package routes

import (
	"log"
	"net/http"

	"go-rest-api-template/model"
	"go-rest-api-template/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterOrderHandlers(r *gin.RouterGroup, service service.OrderService, logger log.Logger) {
	res := OrderResource{service, logger}

	// Register all routes with the handler functions
	r.POST("/orders", res.PostOrderRoute)
	r.GET("/orders/:id", res.GetOrderRoute)
	r.GET("/validateUUID/:id", res.VerifyEmailUUIDRoute)
}

type OrderResource struct {
	service service.OrderService
	logger  log.Logger
}

func (r OrderResource) PostOrderRoute(ctx *gin.Context) {
	// Parse request body into order payload model
	r.logger.Println("Parsing order payload from request")
	orderPayload := model.Order{}
	if err := ctx.BindJSON(&orderPayload); err != nil {
		r.logger.Println("Error in parsing order payload from request: Err: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Error in parsing request payload. Please check if all required fields are present"})
	}

	// Service call to save order details
	res, err := r.service.SaveOrder(orderPayload)

	// Error handling and return response
	if err.Error != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Error, "message": err.Message})
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func (r OrderResource) GetOrderRoute(ctx *gin.Context) {
	// Parse id from URL
	id := ctx.Param("id")
	r.logger.Println("Getting Order from DB; id = ", id)

	// Validate UUID
	orderId, parseErr := uuid.Parse(id)
	if parseErr != nil {
		r.logger.Println("Bad request - Not a valid UUID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": parseErr.Error()})
	}

	// Service call to fetch by id
	res, err := r.service.GetOrderById(orderId)

	// Error handling and return response
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}

func (r OrderResource) VerifyEmailUUIDRoute(ctx *gin.Context) {
	// Parse id from URL
	id := ctx.Param("id")
	r.logger.Println("Verifying Email UUID in DB")

	// Validate UUID
	emailUUID, parseErr := uuid.Parse(id)
	if parseErr != nil {
		r.logger.Println("Bad request - Not a valid UUID")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": parseErr.Error()})
	}

	// Service call to verify UUID
	res, err := r.service.VerifyOrderEmailIsActive(emailUUID)

	// Error handling and return response
	if err != nil {
		if err.Error() == "bad_request" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		ctx.JSON(http.StatusOK, res)
	}
}
