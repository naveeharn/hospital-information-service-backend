package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	_ "github.com/jackc/pgx/v5/stdlib"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

func SetupDatabaseConnection() *sql.DB {
	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		panic("Failed to load .env file")
	}
	// os.Setenv("POSTGRES_USERNAME", "go_trello_01@postgresql.com")
	// os.Setenv("POSTGRES_PASSWORD", "go_trello_01")
	// os.Setenv("POSTGRES_HOSTNAME", "127.0.0.1")
	// os.Setenv("POSTGRES_PORT", "4012")
	// os.Setenv("POSTGRES_DB_NAME", "trello")
	// os.Setenv("JWT_SECRET_KEY", "JWTtrelloSecret220105")
	// defer os.Unsetenv("POSTGRES_USERNAME")
	// defer os.Unsetenv("POSTGRES_PASSWORD")
	// defer os.Unsetenv("POSTGRES_HOSTNAME")
	// defer os.Unsetenv("POSTGRES_PORT")
	// defer os.Unsetenv("POSTGRES_DB_NAME")
	// defer os.Unsetenv("JWT_SECRET_KEY")

	pg_username := os.Getenv("POSTGRES_USERNAME")
	pg_password := os.Getenv("POSTGRES_PASSWORD")
	pg_hostname := os.Getenv("POSTGRES_HOSTNAME")
	pg_port := os.Getenv("POSTGRES_PORT")
	pg_db_name := os.Getenv("POSTGRES_DB_NAME")

	postgres_url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		pg_username, pg_password, pg_hostname, pg_port, pg_db_name,
	)

	// db, err := gorm.Open(postgres.Open(postgres_url), &gorm.Config{})
	// if err != nil {
	// 	panic("Failed to create a connection to postgres database")
	// }

	db, err := sql.Open("pgx", postgres_url)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	// err = db.AutoMigrate(
	// 	&entity.User{},
	// 	&entity.Team{},
	// 	&entity.Dashboard{},
	// 	&entity.Note{},
	// )

	if err != nil {
		panic("Failed to migrate from entities to postgres database")
	}
	log.Println("Succeeded to connect postgres database")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgres, err := db.DB()
	if err != nil {
		panic("Failed to close connection from postgres database")
	}
	dbPostgres.Close()
}
