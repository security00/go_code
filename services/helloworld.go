package services

import (
	context "context"
)

type HelloworldService struct {
}

func (*HelloworldService) SayHello(ctx context.Context, in *RequestMessage) (*ResponseMessage, error) {
	return &ResponseMessage{Name: "宋功名哥哥"}, nil
}
