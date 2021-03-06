package data

import (
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// CreateSession for an existing user.
func (user *User) CreateSession() (session Session, err error) {

	query := `INSERT INTO sessions(uuid, email, user,created_at)
				  VALUES ($1, $2,$3, $4)
				  RETURNING id, uuid, email, user_id, created_at`

	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		createUUID(),
		user.Email,
		user.Id,
		time.Now(),
	).Scan(
		&session.Id,
		&session.Uuid,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)
	return
}

// Session gets the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow(`SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id=$1`, session.UserId).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.Email, &session.CreatedAt)
	return
}

// Check if the database session is valid
func (session *Session) Check() (valid bool, err error) {

	err = Db.QueryRow(`SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id=$1`, session.UserId).
		Scan(&session.Id, &session.Uuid, &session.UserId, &session.Email, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

// DeleteByUUID delete a specific session from the database
func (session *Session) DeleteByUUID() (err error) {
	query := "DELETE FROM sessions WHERE uuid=$1"
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return
}

// User info from the session
func (session *Session)  User() (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT id, uuid, name, email, created_at FROM users WHERE id=$1`, session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// SessionDeleteAll delete all sessions from the database
func SessionDeleteAll() (err error) {
	query := "DELETE FROM sessions"
	_, err = Db.Exec(query)
	return
}

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	query := `INSERT INTO users (uuid, name, email, password, created_at) 
				VALUES ($1, $2, $3, $4, $5)
				RETURNING id, uuid, created_at`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	query := `DELETE FROM users WHERE id=$1`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

// Update update user information in the database
func (user *User) Update() (err error) {
	query := `UPDATE users SET name=$2, email=$3 WHERE id=$1`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

// UserDeleteAll will delete all users
func UserDeleteAll() (err error) {
	query := "DELETE FROM users"
	_, err = Db.Exec(query)
	return
}


// Users Get all users in the database
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// UserByEmail get a user by their email address
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT id, uuid, name, email, password, created_at FROM WHERE email=$1`, email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// UserByUUID get a user by their uuid
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(`SELECT id, uuid, name, email, password, created_at FROM WHERE uuid=$1`, uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

