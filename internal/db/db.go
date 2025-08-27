package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	var err error

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var db *sql.DB

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_ssl_mode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db_host, db_port, db_user, db_password, db_name, db_ssl_mode)

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		fmt.Println("Database not ready, retrying...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Error while connecting to db:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database not responding", err)
	}
	return db
}

func CreateUser(conn *sql.DB, username, passwordHash string) error {
	_, err := conn.Exec("INSERT INTO users(username, password_hash) VALUES ($1, $2)", username, passwordHash)

	if err != nil {
		log.Println("Error inserting user:", err)
	}
	return err
}

func GetPasswordHash(conn *sql.DB, username string) (string, error) {
	var passwordHash string
	err := conn.QueryRow("SELECT password_hash FROM users WHERE username=$1", username).Scan(&passwordHash)
	if err != nil {
		return "", err
	}
	return passwordHash, nil
}

func RunMigrations() {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_ssl_mode := os.Getenv("SSL_MODE")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db_user, db_password, db_host, db_port, db_name, db_ssl_mode,
	)

	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
