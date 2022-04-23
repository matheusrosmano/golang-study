package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectaDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Falha ao carregar o arquivo dotenv")
	}
	conexao := "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@localhost:5432/alura?sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}
	return db
}