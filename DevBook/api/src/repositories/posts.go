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

// UpdatePost updates an existing post
func (repo *Posts) UpdatePost(postID uint64, postToUpdate models.Post) error {
	statement, error := repo.db.Prepare("update posts set title = ?, content = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postToUpdate.Title, postToUpdate.Content, postID); error != nil {
		return error
	}
	return nil
}

// FetchPostByUserID fetch posts of a specific user
func (repo *Posts) FetchPostByUserID(userID uint64) ([]models.Post, error) {
	rows, error := repo.db.Query(`
		select p.id, p.title, p.content, p.likes, p.created_at, p.author_id, u.nick from posts p 
		join users u on p.author_id = u.id where p.author_id = ?
	`, userID)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	// #IMPORTANT non nil slice (length 0)
	userPosts := []models.Post{}
	for rows.Next() {
		var userPost models.Post
		if error = rows.Scan(
			&userPost.ID,
			&userPost.Title,
			&userPost.Content,
			&userPost.Likes,
			&userPost.CreatedAt,
			&userPost.AuthorID,
			&userPost.AuthorNick,
		); error != nil {
			return nil, error
		}
		userPosts = append(userPosts, userPost)
	}
	return userPosts, nil
}

// DeletePost deletes a post from database
func (repo *Posts) DeletePost(postID uint64) error {
	statement, error := repo.db.Prepare("delete from posts where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}

// LikePost adds a like to a post
func (repo *Posts) LikePost(postID uint64) error {
	statement, error := repo.db.Prepare("update posts set likes = likes + 1 where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}

// UnlikePost adds a like to a post
func (repo *Posts) UnlikePost(postID uint64) error {
	statement, error := repo.db.Prepare(`
		update posts set likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		where id = ?
	`)
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(postID); error != nil {
		return error
	}
	return nil
}
