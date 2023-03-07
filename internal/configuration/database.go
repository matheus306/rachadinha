package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const PostgresDriver = "postgres"

var user = "root"

var host = "postgres.cl81fj5g0id8.us-east-2.rds.amazonaws.com"

var password = "123"

const port = "5432"

const dbName = "rachadinha"

const tableName = "tb_grupo"

func Conectar() gorm.DB {

	var secret = recuperarSecret()
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, secret.Username, secret.Password, dbName)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return *db
}
