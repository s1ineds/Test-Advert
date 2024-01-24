package models

type Advertisement struct {
	Id               int
	Category         string
	Title            string
	ShortDescription string
	Description      string
	Image            string
	Price            float32
	Contact          string
	PubDate          string
	CategoryId       string
}
