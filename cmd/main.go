package main

import (
	"context"
	"encoding/json"
	config "rachadinha/internal/configuration"
	"rachadinha/internal/entity"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Request = events.APIGatewayProxyRequest
type Response = events.APIGatewayProxyResponse

func handler(ctx context.Context, r Request) (Response, error) {

	db := config.Conectar()
	var grupo []entity.Grupo

	db.Preload("Usuario").Find(&grupo)

	json, _ := json.Marshal(grupo)
	var resp Response
	resp.Body = string(json)
	resp.StatusCode = 200
	resp.Headers = map[string]string{"Content-Type": "application/json"}

	return resp, nil
}

func main() {
	lambda.Start(handler)
}
