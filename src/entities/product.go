package entities

import "time"

type Student struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duedate     time.Time `json:"duedate"`
	Status      string    `json:"status"`
}
