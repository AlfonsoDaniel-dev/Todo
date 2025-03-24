package DTO

import "time"

type GetTask struct {
	Title     string
	Body      string
	CreatedAt time.Time
}

type CreateTask struct {
	Title string
	Body  string
}
