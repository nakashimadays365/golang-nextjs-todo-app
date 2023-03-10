package model

import "time"

type DBData struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}
