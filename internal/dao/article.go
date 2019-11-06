package dao

import (
	"errors"
	"html/template"
)

type Article struct {
	helper SQLHelper

	Title   string        `json:"name"`
	Brief   string        `json:"brief"`
	Content template.HTML `json:"content"`
	Author  string        `json:"author"`
	Time    int64
}

func (helper SQLHelper) GetArticle(author string, title string) *Article {
	return &Article{
		helper: helper,
		Author: author,
		Title:  title,
	}
}

func (helper SQLHelper) TraversingArticles() ([]Article, error) {
	var articles []Article
	//cmd := "SELECT id, name, author, brief, content FROM article ORDER BY MODIFYTIME DESC"
	cmd := "SELECT title, author, brief, content, timestamp FROM article ORDER BY timestamp DESC"
	rows, err := helper.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		article := &Article{
			helper: helper,
		}
		if err = rows.Scan(&article.Title, &article.Author, &article.Brief, &article.Content, &article.Time); err != nil {
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

func (article *Article) Get() error {
	count := 0
	cmd := "SELECT brief, content, timestamp FROM article WHERE author = ? AND title = ?"
	rows, err := article.helper.db.Query(cmd, article.Author, article.Title)
	if err != nil {
		return err
	}
	for rows.Next() {
		count++
		if err = rows.Scan(&article.Brief, &article.Content, &article.Time); err != nil {
			return err
		}
	}
	if count == 0 {
		return errors.New("Did not find the data with " + article.Author + "'s" + article.Title + ".")
	}
	return err
}

func (article *Article) Insert() error {
	cmd := "INSERT INTO article(title, author, brief, content, timestamp ) VALUES(?, ?, ?, ?, ?)"
	_, err := article.helper.db.Exec(cmd, article.Title, article.Author, article.Brief, article.Content, article.Time)
	return err
}
