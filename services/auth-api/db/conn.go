package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `envconfig:"DB_HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	DBName   string `envconfig:"DB_NAME"`
}

func ConnectDB() (*sql.DB, error) {
	var cfg Config
	envErr := envconfig.Process("", &cfg)
	if envErr != nil {
		log.Fatal(envErr.Error())
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)

	var err error
	for i := range 10 {
		db, err := sql.Open("postgres", psqlInfo)
		if err == nil {
			if err = db.Ping(); err == nil {
				fmt.Println("Connected to " + cfg.DBName)
				return db, nil
			}
		}

		duration := i + 1
		fmt.Println("Database not ready yet... retrying in " + strconv.Itoa(duration) + " seconds")
		time.Sleep(time.Duration(duration) * time.Second)
	}

	log.Fatal("Could not connect to the database after multiple retries.")
	return nil, err
}
