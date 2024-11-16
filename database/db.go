package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	// String de conexão com o Supabase
	stringDeConexao := "host=aws-0-sa-east-1.pooler.supabase.com user=postgres.jkhvplmlxwsxrwycupwf password=0lTaIsIsuPU2Y0gZ dbname=postgres port=6543"

	// Abrir a conexão com o banco de dados
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{
		// Configurações adicionais, se necessário
	})

	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}

	// Configuração do pool de conexões
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter o DB do GORM: %v", err)
	}

	// Definindo o número máximo de conexões abertas
	sqlDB.SetMaxOpenConns(10) // Número máximo de conexões abertas com o banco
	// Definindo o número máximo de conexões ociosas
	sqlDB.SetMaxIdleConns(5) // Número de conexões ociosas
	// Definindo o tempo máximo de vida das conexões
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Tempo máximo de vida de cada conexão

	log.Println("Conectado ao banco de dados com sucesso!")
}
