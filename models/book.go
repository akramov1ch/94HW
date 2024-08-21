package models

type Book struct {
	ID            int    `db:"id" json:"id"`
	Title         string `db:"title" json:"title"`
	Author        string `db:"author" json:"author"`
	PublishedDate string `db:"published_date" json:"published_date"`
	ISBN          string `db:"isbn" json:"isbn"`
}
