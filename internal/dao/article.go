package dao

type Article struct {
	helper SQLHelper

	Id 		string `json:"id"`
	Name 	string `json:"name"`
	Brief 	string `json:"brief"`
	Content string `json:"content"`
	Author 	string `json:"author"`
}

func (helper SQLHelper) GetArticle(id string) *Article {
	return &Article{
		helper:  helper,
		Id:      id,
	}
}

func (article *Article) Get() error {
	cmd := "SELECT name, author, brief, content FROM article WHERE id = ?"
	rows, err := article.helper.db.Query(cmd, article.Id)
	if err != nil {
		return err
	}
	for rows.Next() {
		if err = rows.Scan(&article.Name, &article.Author, &article.Brief, &article.Content); err != nil {
			return err
		}
	}
	return err
}

func (article *Article) Insert() error {
	cmd := "INSERT INTO article(id, name, author, brief, content) VALUES(?, ?, ?, ?, ?)"
	_, err := article.helper.db.Exec(cmd, article.Id, article.Name, article.Author, article.Brief, article.Content)
	return err
}