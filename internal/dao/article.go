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
	cmd := "SELECT "
}