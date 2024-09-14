package types

type SuccessResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

type SuccessDeleteResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	
}

