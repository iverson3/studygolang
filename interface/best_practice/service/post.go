package service

import "studygolang/interface/best_practice/model"

type Service interface {
	ListPosts() ([]*model.Post, error)
}

type service struct {
	conn string
}

func (s *service) ListPosts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	posts = append(posts, &model.Post{Name: "name1"})
	return posts, nil
}

func NewService(conn string) Service {
	return &service{conn: conn}
}
