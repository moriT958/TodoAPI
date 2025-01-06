package store

import (
	"context"
	"database/sql"
	"io"
	"log"
	"os"
	"testing"
)

var (
	testDsn     = os.Getenv("DATABASE_URL")
	testDB      *sql.DB
	testFixture *sql.Tx
)

var mockStore *TodoStore

func TestMain(m *testing.M) {
	if err := setupDB(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func setupDB() error {
	var err error
	ctx := context.Background()

	testDB, err = sql.Open("postgres", testDsn)
	if err != nil {
		return err
	}

	testFixture, err = testDB.Begin()
	if err != nil {
		return err
	}

	mockStore = NewTodoStore(testFixture)

	// read setup sql file
	file, err := os.Open("testdata/setup.sql")
	if err != nil {
		return err
	}
	defer file.Close()

	sqlBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	sqlStr := string(sqlBytes)

	if _, err := mockStore.db.ExecContext(ctx, sqlStr); err != nil {
		return err
	}

	return nil
}

func teardown() {
	testDB.Close()
	testFixture.Rollback()
}
