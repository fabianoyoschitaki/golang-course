package repositories

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB // will receive a database from controllers
}

// NewPostsRepository creates Posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	// inside this struct we'll have the database operations, insert, update etc.
	// #IMPORTANT: controller only opens connection, repository makes connection with tables
	return &Posts{db}
}

// Create creates a new post
func (repo Posts) Create(newPost models.Post) (uint64, error) {
	statement, error := repo.db.Prepare(
		"insert into posts (title, content, author_id) values (?, ?, ?)",
	)
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(newPost.Title, newPost.Content, newPost.AuthorID)
	if error != nil {
		return 0, error
	}

	// at this point the post is already inserted into database
	lastInsertedPostID, error := result.LastInsertId()
	if error != nil {
		return 0, nil
	}
	return uint64(lastInsertedPostID), nil
}

// FetchPostByID returns a post by its ID
func (repo *Posts) FetchPostByID(postID uint64) (models.Post, error) {
	rows, error := repo.db.Query(`
		select p.id, p.title, p.content, p.author_id, p.likes, p.created_at, u.nick 
		from posts p inner join users u
		on u.id = p.author_id where p.id = ?`, postID)
	if error != nil {
		return models.Post{}, error
	}
	defer rows.Close()

	var post models.Post
	if rows.Next() {
		if error = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return models.Post{}, error
		}
	}
	return post, nil
}

// FetchPosts returns slice of posts
func (repo *Posts) FetchPosts(userID uint64) ([]models.Post, error) {
	rows, error := repo.db.Query(`
		select distinct p.id, p.title, p.content, p.author_id, p.likes, p.created_at, u.nick 
		from posts p
		inner join users u on u.id = p.author_id
		inner join user_followers uf on p.author_id = uf.user_id
		where u.id = ? or uf.follower_user_id = ?
		order by p.id desc
	`, userID, userID)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if error = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); error != nil {
			return nil, error
		}
		posts = append(posts, post)
	}
	return posts, nil
}
