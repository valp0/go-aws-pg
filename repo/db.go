package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/lib/pq"
)

const (
	dbName = "postgres"
	dbUser = "iampg"
	dbHost = "rds-postgresql-test.clmnif2sowhb.us-east-1.rds.amazonaws.com"
	dbPort = 5432
	region = "us-east-1"
)

// Creates a connection to a PostgreSQL database using the global constants.
func connectDB() (*sql.DB, error) {
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
