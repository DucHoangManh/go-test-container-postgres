package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go-test-container-postgres/internal/persistence"
)

var postgresContainer testcontainers.Container

func StartDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:12",
		ExposedPorts: []string{"5432/tcp"},
		HostConfigModifier: func(config *container.HostConfig) {
			config.AutoRemove = true
		},
		Env: map[string]string{
			"POSTGRES_USER":     "denishoang",
			"POSTGRES_PASSWORD": "pgpassword",
			"POSTGRES_DB":       "products",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgres, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		os.Exit(1)
	}
	postgresContainer = postgres
}

func NewDatabase(t *testing.T) *sqlx.DB {
	ctx := context.Background()
	if postgresContainer == nil {
		t.Fatal("postgres is not yet started")
	}
	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal("err get mapped port from container")
	}

	hostIP, err := postgresContainer.Host(ctx)
	// open connection to postgres instance in order to create other databases
	baseDb, err := sqlx.Open(
		"postgres", fmt.Sprintf(
			"postgres://denishoang:pgpassword@%s:%s/products?sslmode=disable", hostIP, mappedPort.Port(),
		),
	)
	defer func() {
		if err := baseDb.Close(); err != nil {
			t.Fatal("err close connection to db")
		}
	}()
	dbName := fmt.Sprintf("%s_%d", "products", rand.Int63())
	if _, err := baseDb.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)); err != nil {
		t.Fatal("err creating postgres database")
	}
	connString := fmt.Sprintf(
		"postgres://denishoang:pgpassword@%s:%s/%s?sslmode=disable", hostIP, mappedPort.Port(), dbName,
	)
	// apply migrations to the newly created database
	if err = persistence.MigrationUp(connString); err != nil {
		t.Fatal("err connect postgres database")
	}
	// connect to new database
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		t.Fatal("err connect postgres database")
	}
	return db
}
