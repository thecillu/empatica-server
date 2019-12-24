package model

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

var DB_CONNECTION_STRING string

func init() {
	var DB_ENDPOINT string
	DB_ENDPOINT = os.Getenv("DB_ENDPOINT")
	var DB_USERNAME string
	DB_USERNAME = os.Getenv("DB_USERNAME")
	var DB_PASSWORD string
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_CONNECTION_STRING = DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_ENDPOINT + ")/empatica"
}

type Article struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Content     string `json:"Content"`
}

func GetAllArticles() ([]Article, error) {
	var articles []Article
	db, err := sql.Open("mysql", DB_CONNECTION_STRING)
	if err != nil {
		return nil, err
	}

	results, err := db.Query("SELECT ID, title, description, content FROM articles")
	if err != nil {
		fmt.Println(err)
		return articles, err
	}

	for results.Next() {
		article := Article{}
		err = results.Scan(&article.ID, &article.Title, &article.Description, &article.Content)
		articles = append(articles, article)
	}
	if articles == nil {
		articles = []Article{}
	}
	return articles, nil

}

func SaveArticle(article Article) error {
	db, err := sql.Open("mysql", DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println(article.ID)
	sqlStatement := `
	INSERT INTO articles (id, title, description, content) 
	VALUES (?, ?, ?, ?)`
	_, err = db.Exec(sqlStatement, article.ID, article.Title, article.Description, article.Content)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}

func UpdateArticle(article Article) (int, error) {
	var errExistentArticle error
	_, errExistentArticle = GetArticle(article.ID)
	db, err := sql.Open("mysql", DB_CONNECTION_STRING)
	if err != nil {
		return 0, err
	}
	defer db.Close()
	if errExistentArticle != nil {
		err = SaveArticle(article)
		if err != nil {
			return 500, err
		} else {
			return http.StatusCreated, nil
		}
	} else {
		sqlStatement := `
		UPDATE articles set title = ?,  description = ?,  content = ?
		WHERE ID = ?`
		_, err = db.Exec(sqlStatement, article.Title, article.Description, article.Content, article.ID)
		if err != nil {
			fmt.Println(err)
			return 500, err
		} else {
			return http.StatusOK, nil
		}
	}
}

func GetArticle(id string) (Article, error) {
	var article Article
	db, err := sql.Open("mysql", DB_CONNECTION_STRING)
	defer db.Close()
	err = db.QueryRow("SELECT ID, title, description, content FROM articles where id = ?", id).Scan(&article.ID, &article.Title, &article.Description, &article.Content)
	if err != nil {
		fmt.Println(err)
		return article, err
	} else {
		return article, nil
	}
}

func DeleteArticle(id string) error {
	db, err := sql.Open("mysql", DB_CONNECTION_STRING)
	if err != nil {
		return err
	}
	defer db.Close()
	sqlStatement := `
	DELETE FROM articles
	WHERE ID = ?`
	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}
