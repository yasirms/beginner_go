package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := db_string()
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")

	createEmployee := `CREATE TABLE IF NOT EXISTS employee (
	id SERIAL PRIMARY KEY,
	employee_id VARCHAR(50) NOT NULL,
	name VARCHAR(50) NOT NULL,
	father_name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL UNIQUE,
	address VARCHAR(50) NOT NULL,
	phone_number VARCHAR(50) NOT NULL,
	cnic_number VARCHAR(50) NOT NULL,
	job_title VARCHAR(50) NOT NULL,
	job_start_date VARCHAR(50) NOT NULL,
	job_end_date VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`
	_, err = DB.Exec(createEmployee)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employees Table created successfully")

	if err := runMigrations(DB); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully")
}

func db_string() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv("POSTGRES_DB")
}



func runMigrations(db *sql.DB) error{
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create driver: %w", err)	
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {	
		return fmt.Errorf("up migrations failed: %w", err)
	}
	return nil
}