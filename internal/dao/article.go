package dao

import "errors"

type Article struct {
	helper SQLHelper

	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Brief 	string `json:"brief"`
	Content string `json:"content"`
	Author 	string `json:"author"`
	Time 	int64
}

func (helper SQLHelper) GetArticle(id string) *Article {
	return &Article{
		helper:  helper,
		Id:      id,
	}
}

func (helper SQLHelper) TraversingArticles() ([]Article, error) {
	var articles []Article
	//cmd := "SELECT id, name, author, brief, content FROM article ORDER BY MODIFYTIME DESC"
	cmd := "SELECT id, name, author, brief, content, timestamp FROM article ORDER BY timestamp DESC"
	rows, err := helper.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		article := &Article{
			helper:  helper,
		}
		if err = rows.Scan(&article.Id, &article.Name, &article.Author, &article.Brief, &article.Content, &article.Time); err != nil {
			return nil, err
		}
		articles = append(articles, *article)
	}
	return articles, nil
}

func (article *Article) Get() error {
	count := 0
	cmd := "SELECT name, author, brief, content, timestamp FROM article WHERE id = ?"
	rows, err := article.helper.db.Query(cmd, article.Id)
	if err != nil {
		return err
	}
	for rows.Next() {
		count++
		if err = rows.Scan(&article.Name, &article.Author, &article.Brief, &article.Content, &article.Time); err != nil {
			return err
		}
	}
	if count == 0 {
		return errors.New("Did not find the data with id " + article.Id + ".")
	}
	return err
}

func (article *Article) Insert() error {
	cmd := "INSERT INTO article(id, name, author, brief, content, timestamp ) VALUES(?, ?, ?, ?, ?, ?)"
	_, err := article.helper.db.Exec(cmd, article.Id, article.Name, article.Author, article.Brief, article.Content, article.Time)
	return err
}