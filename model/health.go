package model

import "time"

type Health struct {
	ServiceName   string    `json:"serviceName"`
	HealthStatus  string    `json:"healthStatus"`
	ServiceStatus string    `json:"serviceStatus"`
	CurrentTime   time.Time `json:"currentTime"`
}

type DefaultResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
