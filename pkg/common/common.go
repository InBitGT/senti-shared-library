package common

// Estructura base para todas las respuestas
type APIResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Información del error
type ErrorInfo struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Data    *string `json:"data,omitempty"`
}

// Respuesta exitosa
func SuccessResponse(message string, data interface{}, code string) APIResponse {
	return APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// Respuesta de error
func ErrorResponse(statusCode int, code string, message string, data *string) ErrorInfo {
	return ErrorInfo{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (e *ErrorInfo) Error() string {
	return e.Message
}
