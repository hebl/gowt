package core

import (
	"time"
)

//======================================
// 模型

type User struct {
	Id    int
	Name  string
	Email string
	Dtime time.Time
}

type Users []*User

func (users *Users) New() interface{} {
	u := &User{}
	*users = append(*users, u)
	return u
}

type Post struct {
	Id      int
	Title   string
	Content string
	Created time.Time
}

type Posts []*Post

func (posts *Posts) New() interface{} {
	p := &Post{}
	*posts = append(*posts, p)
	return p
}

//======================================
// 数据库操作
func createUser(user *User) error {
	q := `INSERT INTO users VALUES(?id, ?name, ?email, ?dtime)`
	_, err := db.ExecOne(q, user)

	return err
}

func updateUser(user *User) error {
	q := `UPDATE users SET name = ?name, email = ?email, dtime=?dtime WHERE id = ?id`
	_, err := db.ExecOne(q, user)

	return err
}

func userById(id int) (*User, error) {
	u := &User{}
	q := `SELECT * FROM users WHERE id = ?`
	_, err := db.QueryOne(u, q, id)

	if err != nil {
		return nil, err
	}

	return u, nil

}

func queryUsers() ([]*User, error) {
	q := `SELECT * FROM users`
	var users Users
	_, err := db.Query(&users, q)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func createPost(post *Post) error {
	q := `INSERT INTO post VALUES(?id, ?title, ?content, now())`
	_, err := db.ExecOne(q, post)

	return err
}

func updatePost(post *Post) error {
	q := `UPDATE post SET title = ?title, content = ?content, created=?created WHERE id = ?id`
	_, err := db.ExecOne(q, post)

	return err
}

func postById(id int) (*Post, error) {
	p := &Post{}
	q := `SELECT * FROM post WHERE id = ?`
	_, err := db.QueryOne(p, q, id)

	if err != nil {
		return nil, err
	}

	return p, nil

}

func queryPosts() ([]*Post, error) {
	q := `SELECT * FROM post`
	var posts Posts
	_, err := db.Query(&posts, q)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
