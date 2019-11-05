package utilities

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	u "github.com/idirall22/user"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

var userUsernameTest = "alice"
var userPasswordTest = "fdpjfd654/*sMLdf"

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

// LoginUser log a user and get a JWT token to test with
func LoginUser(db *sql.DB) string {

	m := make(map[string]string)
	m["username"] = userUsernameTest
	m["password"] = userPasswordTest

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	body := bytes.NewReader(b)

	serviceUser := u.StartService(db, "users")
	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/login", body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	h := http.HandlerFunc(serviceUser.Login)
	h.ServeHTTP(w, r)

	return w.Header().Get("Authorization")
}
