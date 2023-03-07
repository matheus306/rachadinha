package config

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Secret struct {
	Username string
	Password string
}

func recuperarSecret() Secret {

	secretName := "rds!db-1d34b433-388a-4692-8bde-b835ee14b645"
	region := "us-east-2"
	sess := session.Must(session.NewSession())

	svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(region))

	result, err := svc.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
	if err != nil {
		log.Fatal(err.Error())
	}

	var secretValue Secret

	json.Unmarshal([]byte(*result.SecretString), &secretValue)

	return secretValue
}
