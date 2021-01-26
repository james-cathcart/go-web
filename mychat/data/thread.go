package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        string
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// format the CreateAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	// Format Thread date
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	// Format Post date
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// NumReplies Get the number of posts for a
func (thread *Thread) NumReplies() (count int, err error) {
	rows, err := Db.Query(`SELECT count(*) FROM posts WHERE thread_id=$1`, thread.Id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	return
}

// Posts get all posts for a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err :=
		Db.Query(
			`SELECT id, uuid, body, user_id, thread_id, created_at 
					FROM posts 
					WHERE thread_id=$1`,
			thread.Id,
		)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	return
}

// CreateThread creates a new Thread
func (user *User) CreateThread(topic string) (thread Thread, err error) {
	query := `INSERT INTO threads (uuid, topic, user_id, created_at) 
				VALUES ($1, $2, $3, $4) 
				RETURNING id, uuid, topic, user_id, created_at`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	// Use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(
		createUUID(),
		topic,
		user.Id,
		time.Now(),
	).Scan(
		&thread.Id,
		&thread.UserId,
		&thread.Topic,
		&thread.UserId,
		&thread.CreatedAt,
	)

	return
}

// CreatePost add a post to an existing thread
func (user *User) CreatePost(thread Thread, body string) (post Post, err error) {
	query := `INSERT INTO posts (uuid, body, user_id, thread_id, created_at) 
				VALUES ($1, $2, $3 $4, $5) 
				RETURNING id, uuid, body, user_id, thread_id, created_at`
	stmt, err := Db.Prepare(query)
	if err != nil {
		return
	}
	defer stmt.Close()

	stmt.QueryRow(
		createUUID(),
		body,
		thread.UserId,
		thread.Id,
		thread.CreatedAt,
	).Scan(
		&post.Id,
		&post.Uuid,
		&post.Body,
		&post.UserId,
		&post.ThreadId,
		&post.CreatedAt,
	)
	return
}

// Threads return all threads from the database
func Threads() (threads []Thread, err error) {

	rows, err := Db.Query(
		`SELECT id, uuid, topic, user_id, creation_at 
				FROM threads 
				ORDER BY created_at 
				DESC`)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		thread := Thread{}

		err = rows.Scan(
			&thread.Id,
			&thread.Uuid,
			&thread.Topic,
			&thread.UserId,
			&thread.CreatedAt,
		)
		if err != nil {
			return
		}

		threads = append(threads, thread)
	}
	return
}

// ThreadByUUID get a thread by the UUID
func ThreadByUUID(uuid string) (thread Thread, err error) {
	thread = Thread{}
	err = Db.QueryRow(
		`SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid=$1`, uuid,
	).Scan(
		&thread.Id,
		&thread.Uuid,
		&thread.Topic,
		&thread.UserId,
		&thread.CreatedAt,
	)

	return
}

// User get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow(
		`SELECT id, uuid, topic, user_id, created_at FROM users WHERE id=$1`, thread.UserId,
	).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// User get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow(
		`SELECT id, uuid, topic, user_id, created_at FROM users WHERE id=$1`, post.UserId,
	).Scan(
		&user.Id,
		&user.Uuid,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	return
}
