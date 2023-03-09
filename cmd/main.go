package main

import (
	"context"
	"encoding/json"

	"rachadinha/internal/entity"
	service "rachadinha/internal/middleware"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/lib/pq"
)

type Request = events.APIGatewayProxyRequest
type Response = events.APIGatewayProxyResponse

func handler(ctx context.Context, r Request) (Response, error) {

	var grupos []entity.Grupo
	service.RecuperarGrupos(&grupos)

	var usuarioDoMes entity.Usuario
	var proximoSequencial int

	for _, grupo := range grupos {
		usuarioDoMes, proximoSequencial = service.RecuperarUsuarioDoMes(&grupo)
		grupo.SquencialAtual = proximoSequencial
		service.SalvarGrupo(&grupo)
	}

	json, _ := json.Marshal(usuarioDoMes)
	var resp Response
	resp.Body = string(json)
	resp.StatusCode = 200
	resp.Headers = map[string]string{"Content-Type": "application/json"}

	return resp, nil
}

func main() {
	lambda.Start(handler)
}
