package response

type ResponseDataDTO struct {
	StatusCode int `json:"status_code"`
	Data       any `json:"data"`
}
