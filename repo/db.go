package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/lib/pq"
)

// Creates a connection to a PostgreSQL database using the global constants.
func connectDB() (*sql.DB, error) {

	var (
		dbName string = os.Getenv("AWS_DB_NAME")
		dbUser string = os.Getenv("AWS_DB_USER")
		dbHost string = os.Getenv("AWS_DB_HOST")
		port   string = os.Getenv("AWS_DB_PORT")
		region string = os.Getenv("AWS_REGION")
	)

	dbPort, _ := strconv.Atoi(port)

	var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to create authentication token: " + err.Error())
	}

	dbConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		dbHost, dbPort, dbUser, authenticationToken, dbName,
	)

	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare connection to database %s, %s", dbName, err)
	}

	return db, nil
}

// Closes open connection with db.
func (r repository) CloseDB() error {
	return r.db.Close()
}
