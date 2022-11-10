package internal

import (
	"errors"
	"strings"
)

type Response struct {
	IP string `json:"ip"`
}

func parse(ip string) (Response, error) {
	value := strings.TrimSpace(ip)
	if value == "" {
		return Response{}, errors.New("invalid parameter(s) supplied")
	}
	parts := strings.Split(value, ":")
	if len(parts) < 1 {
		return Response{}, errors.New("invalid parameter(s) supplied")
	} else {
		return Response{
			IP: parts[0],
		}, nil
	}
}
