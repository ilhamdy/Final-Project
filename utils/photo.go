package utils

import (
	"time"
)

type GetPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title,"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	User      *User      `json:"user"`
}

type ResponseDataGetPhoto struct {
	Status string     `json:"status" example:"success"`
	Data   []GetPhoto `json:"data"`
}

type CreatePhoto struct {
	Title    string `json:"title" example:"A Title"`
	Caption  string `json:"caption" example:"A caption"`
	PhotoUrl string `json:"photo_url" example:"https://www.example.com/image.jpg"`
}

type CreatedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	CreatedAt *time.Time `json:"created_at"`
}

type ResponseDataCreatedPhoto struct {
	Status string       `json:"status" example:"success"`
	Data   CreatedPhoto `json:"data"`
}

type UpdatePhoto struct {
	Title    string `json:"title" example:"A new title"`
	Caption  string `json:"caption" example:"A new caption"`
	PhotoUrl string `json:"photo_url" example:"https://www.example.com/new-image.jpg"`
}

type UpdatedPhoto struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Caption   string     `json:"caption"`
	PhotoUrl  string     `json:"photo_url"`
	UserID    string     `json:"user_id"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type ResponseDataUpdatedPhoto struct {
	Status string       `json:"status" example:"success"`
	Data   UpdatedPhoto `json:"data"`
}

type ResponseMessageDeletedPhoto struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"your photo has been successfully deleted"`
}

type ResponseMessagePhoto struct {
	Status string `json:"status" example:"fail"`
	Data   string `json:"data" example:"the error explained here"`
}
