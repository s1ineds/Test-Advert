package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	ConnectionString string
}

func (db Database) connectToDb() *sql.DB {
	if db.ConnectionString == "" {
		log.Fatal("Connection string is empty!")
	}

	conn, err := sql.Open("postgres", db.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func (db Database) GetEntries() []Advertisement {
	connection := db.connectToDb()
	defer connection.Close()

	sql := `SELECT advId, category, title, short_description, description, image, price, contact, pub_date FROM board.category
	INNER JOIN board.advert
	ON board.category.category_id=board.advert.category_id`

	rows, err := connection.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	var listOfAdverts []Advertisement

	for rows.Next() {
		adv := Advertisement{}
		err := rows.Scan(&adv.Id, &adv.Category, &adv.Title, &adv.ShortDescription, &adv.Description,
			&adv.Image, &adv.Price, &adv.Contact, &adv.PubDate)
		if err != nil {
			log.Fatal(err)
		}
		listOfAdverts = append(listOfAdverts, adv)
	}
	return listOfAdverts
}

func (db Database) GetCategories() map[string]string {
	connection := db.connectToDb()
	defer connection.Close()

	sqlQuery := `SELECT category_id, category_name FROM board.category;`
	rows, err := connection.Query(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}

	var catId, tmpCat string
	var categories map[string]string = make(map[string]string, 0)

	for rows.Next() {
		rows.Scan(&catId, &tmpCat)
		categories[catId] = tmpCat
	}

	return categories
}

func (db Database) AddEntries(object Advertisement) {
	defer db.recoverFromPanic()
	connection := db.connectToDb()
	defer connection.Close()

	sqlQuery := `INSERT INTO board.advert(title, short_description, description, image, price, contact, pub_date, category_id) VALUES 
	($1, $2, $3, $4, $5, $6, $7, $8)`

	sqlResult, err := connection.Exec(sqlQuery, object.Title, object.ShortDescription, object.Description,
		object.Image, object.Price, object.Contact, object.PubDate, object.CategoryId)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, _ := sqlResult.RowsAffected()
	fmt.Printf("Rows affected: %d", rowsAffected)
}

func (db Database) GetSpecificAdvert(advId string) *Advertisement {
	defer db.recoverFromPanic()
	connection := db.connectToDb()
	defer connection.Close()

	sqlQuery := `SELECT title, image, description, contact, price FROM board.advert
	WHERE advid=$1`

	rows, err := connection.Query(sqlQuery, advId)
	if err != nil {
		log.Fatal(err)
	}

	var newAdvert Advertisement = Advertisement{}

	for rows.Next() {
		scanErr := rows.Scan(&newAdvert.Title, &newAdvert.Image, &newAdvert.Description, &newAdvert.Contact, &newAdvert.Price)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
	}
	return &newAdvert
}

func (db Database) recoverFromPanic() {
	if result := recover(); result != nil {
		fmt.Println("Recovered. Error:\n", result)
	}
}
