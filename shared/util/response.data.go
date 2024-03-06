package util

import "favorite-book/delivery/http/dto/response"

func ConstructResponseSuccess(statusCode int, dataMsg any) (response.ResponseDataDTO, int) {
	resp := response.ResponseDataDTO{
		StatusCode: statusCode,
		Data:       dataMsg,
	}

	return resp, statusCode
}
