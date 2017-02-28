package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" //github.com/mattn/go-sqlite3
)

const (
	dbName = "mywiki"
)

type SqlDb struct {
	db *sql.DB
}

// Init connects to an external sqlite DB and returns a DB implementation
// that fronts that connection. call Close() on the returned value when done.
func Init() (SqlDb, error) {
	db, err := sql.Open("sqlite3", "./"+dbName+".db")
	if err != nil {
		log.Fatal(err)
	}
	return SqlDb{db: db}, err
}

// Close releases the underlying connections. always call this when
// completely done with operations, not before.
func (m *SqlDb) Close() {
	m.db.Close()
}

// GetArticle gets all the categories associated with a username and
// places it inside Model
// returns the value and nil if it was found, nil and ErrNotFound if it
// was not found and another appropriate error otherwise
func (m SqlDb) GetArticle(username string, title string) (Article, error) {
	query := "SELECT * FROM Article WHERE username = ? AND title = ?"
	rows, err := m.db.Query(query, username, title)

	var article Article

	if err != nil {
		return Article{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&article.Username,
			&article.Title,
			&article.Content,
			&article.DateCreated,
			&article.DateModified,
			&article.Category,
		)
		if err != nil {
			return Article{}, err
		}
	}
	rows.Close()

	return article, err
}

// GetArticlesByUser gets all the articles associated with that user
// returns the value and nil if it was found, nil and ErrNotFound if it
// was not found and another appropriate error otherwise
func (m SqlDb) GetAllArticlesByUser(username string) ([]Article, error) {
	query := "SELECT * FROM Article WHERE username = ?"
	rows, err := m.db.Query(query, username)

	var article Article
	var articles []Article

	if err != nil {
		return []Article{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&article.Username,
			&article.Title,
			&article.Content,
			&article.DateCreated,
			&article.DateModified,
			&article.Category,
		)
		if err != nil {
			return []Article{}, err
		}
		articles = append(articles, article)
	}
	rows.Close()
	return articles, err
}

func (m SqlDb) GetAllArticlesByCategory(category string) ([]Article, error) {
	query := "SELECT * FROM Article WHERE category = ?"
	rows, err := m.db.Query(query, category)

	var article Article
	var articles []Article

	if err != nil {
		return []Article{}, err
	}
	for rows.Next() {
		err = rows.Scan(
			&article.Username,
			&article.Title,
			&article.Content,
			&article.DateCreated,
			&article.DateModified,
			&article.Category,
		)
		if err != nil {
			return []Article{}, err
		}
		articles = append(articles, article)
	}
	rows.Close()
	return articles, err
}

func (m SqlDb) UpsertArticle(article Article) (bool, error) {
	query := "INSERT OR REPLACE INTO Article(username, title, content, created_date, last_modified_at, cat_name) values (?,?,?,?,?,?)"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		log.Println("STATEMENT ERROR")
		log.Print(err)
		return false, err
	}
	res, err := stmt.Exec(
		article.Username,
		article.Title,
		article.Content,
		article.DateCreated,
		article.DateModified,
		article.Category,
	)
	_ = res
	if err != nil {
		log.Println("EXEC ERROR")
		log.Print(err)
		return false, err
	}
	return true, nil
}
