package utilities

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

// ConnectDataBaseTest create a connection to database test
func ConnectDataBaseTest() (*sql.DB, error) {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// BuildDataBase build table to test with
func BuildDataBase(db *sql.DB, query string) error {

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

// CloseDataBaseTest close db
func CloseDataBaseTest(db *sql.DB) {
	db.Close()
}
