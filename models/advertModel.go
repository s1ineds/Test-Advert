package models

// type Advertisement struct {
// 	Id               int
// 	Category         string
// 	Title            string
// 	ShortDescription string
// 	Description      string
// 	Image            string
// 	Price            float32
// 	Contact          string
// 	PubDate          string
// 	CategoryId       string
// }
type Advertisement struct {
	Id               int     `json:"id"`
	Category         string  `json:"category"`
	Title            string  `json:"title"`
	ShortDescription string  `json:"short_description"`
	Description      string  `json:"description"`
	Image            string  `json:"image"`
	Price            float32 `json:"price"`
	Contact          string  `json:"contact"`
	PubDate          string  `json:"pub_date"`
	CategoryId       string  `json:"category_id"`
}
