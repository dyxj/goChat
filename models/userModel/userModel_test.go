package userModel_test

import (
	"database/sql"
	"fmt"
	"goChat/models"
	"goChat/models/userModel"
	"log"
	"os"
	"testing"
)

var usrDB *sql.DB

func TestMain(m *testing.M) {
	fmt.Println("Start Main")
	var err error
	usrDB, err = models.NewDatabase()
	if err != nil {
		log.Fatalf("TestUserModel Failure 1: ", err)
	}
	defer usrDB.Close()
	ensureUserTableCreated(usrDB)
	code := m.Run()

	fmt.Println("End Main")
	os.Exit(code)
}

func TestUserModelFunctions(t *testing.T) {
	u := &userModel.User{}
	u.UserId = "user1"
	u.Password = "password1"
	var err error

	// Insert
	u.UID, err = userModel.Insert(u, usrDB)
	fmt.Println(u.UID)
	if err != nil {
		t.Errorf("Insert failed: %v\n", err)
	}

	// Get
	gUsr, err := userModel.GetOne(u.UID, usrDB)
	if err != nil {
		t.Errorf("Get one user fail %v\n", err)
	}
	if gUsr.UID != u.UID && gUsr.UserId != u.UserId {
		t.Errorf("Get wrong user\n")
	}
	fmt.Println(gUsr)

	// Get list
	gUU, err := userModel.GetList(0, 100,usrDB)
	if err != nil {
		t.Errorf("Get one user fail %v\n", err)
	}
	if len(gUU) < 1 {
		t.Errorf("No users found\n")
	}

	// Delete
	cnt, err := userModel.Delete(u.UID, usrDB)
	if err != nil {
		t.Errorf("Delete failed: %v\n", err)
	}
	if cnt != 1 {
		t.Errorf("No users deleted\n")
	}
}

func ensureUserTableCreated(db *sql.DB) {
	if _, err := db.Exec(userTableCreate); err != nil {
		log.Fatal(err)
	}
}

const userTableCreate = `
CREATE TABLE IF NOT EXISTS USERS
(
	uid uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
	userid text NOT NULL UNIQUE,
	password text NOT NULL,
	first_name text NOT NULL,
	last_name text NOT NULL,
	email text NOT NULL,
	created_date date NOT NULL DEFAULT 'now'::text::date,
	last_updated timestamp NOT NULL DEFAULT now(),
	last_signin timestamp,
	active bool DEFAULT true
)
`
