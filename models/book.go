package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MediumPrice float32   `json:"medium_price"`
	Author      string    `json:"author"`
	ImageUrl    string    `json:"img_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateddAt  time.Time `json:"updated_at"`
}
