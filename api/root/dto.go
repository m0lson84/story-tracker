// Package root provides HTTP handlers and data transfer objects for root API endpoints.
package root

// Response is the response for the root API endpoint.
type Response struct {
	// Response body
	Body struct {
		// The message to return
		Message string `json:"message" example:"Hello world!" doc:"Root API response"`
	}
	// HTTP status code
	Status int `json:"status" example:"200" doc:"HTTP status code"`
}

// HealthResponse is the response for the health check API endpoint.
type HealthResponse struct {
	// Response body
	Body struct {
		// The status of the database connection
		DB map[string]string `json:"db" doc:"Database health check"`
	}
	// HTTP status code
	Status int `json:"status" example:"200" doc:"HTTP status code"`
}
