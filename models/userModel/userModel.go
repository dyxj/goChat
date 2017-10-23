package userModel

import (
	"database/sql"
	"time"
)

type User struct {
	//UID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UID         string    `json:"id" bson:"_id,omitempty"`
	UserId      string    `json:"userid" bson:"userid"`
	Password    string    `json:"password" bson:"password"`
	FirstName   string    `json:"firstname" bson:"firstname"`
	LastName    string    `json:"lastname" bson:"lastname"`
	Email       string    `json:"email" bson:"email"`
	Created     time.Time `json:"created" bson:"created"`
	LastUpdated time.Time `json:"lastupdated" bson:"lastupdated"`
	LastSignIn  time.Time `json:"lastsignin" bson:"lastsignin"`
	Active      bool      `json:"active" bson:"active"`
}

// Insert: Inserts new user to users table in database
func Insert(u *User, db *sql.DB) (string, error) {
	var newUid string
	err := db.QueryRow(`
INSERT INTO users
(userid, password, first_name, last_name, email, created_date, last_updated, last_signin, active)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING uid`,
		u.UserId, u.Password, u.FirstName, u.LastName, u.Email, u.Created,
		u.LastUpdated, u.LastSignIn, u.Active).Scan(&newUid)

	if err != nil {
		return newUid, err
	}

	return newUid, nil
}

// Delete: Deletes user by uid from users table
func Delete(uid string, db *sql.DB) (int64, error) {
	rslt, err := db.Exec(`DELETE FROM users WHERE uid=$1`, uid)
	if err != nil {
		return 0, err
	}
	rCnt, err := rslt.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rCnt, nil
}

// Get: Get user with uid from users table
func GetOne(uid string, db *sql.DB) (User, error) {
	u := User{}
	err := db.QueryRow(`
SELECT uid, userid, password, first_name,
last_name, email, created_date, last_updated, last_signin, active
FROM users
WHERE uid = $1`,
		uid).Scan(&u.UID, &u.UserId, &u.Password, &u.FirstName, &u.LastName,
		&u.Email, &u.Created, &u.LastUpdated, &u.LastSignIn, &u.Active)

	if err != nil {
		return u, err
	}

	return u, nil
}

// GetList: Get list of users from users table
func GetList(start int, count int, db *sql.DB) ([]User, error) {
	rows, err := db.Query(`
SELECT uid, userid, password, first_name,
last_name, email, created_date, last_updated, last_signin, active
FROM users
ORDER BY userid
LIMIT $1 OFFSET $2`,
		count, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uu := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.UID, &u.UserId, &u.Password, &u.FirstName, &u.LastName,
			&u.Email, &u.Created, &u.LastUpdated, &u.LastSignIn, &u.Active)
		if err != nil {
			return nil, err
		}
		uu = append(uu, u)
	}
	return uu, nil
}
