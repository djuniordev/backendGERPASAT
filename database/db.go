package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.jkhvplmlxwsxrwycupwf password=0lTaIsIsuPU2Y0gZ dbname=postgres port=6543 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
