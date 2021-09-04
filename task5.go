package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ad struct
type Ad struct {
	id          uint16 `json:"id"`         // id of ad
	Name        string `json:"name"`       // name element of ad
	Description string `json:"descripion"` // descripion element of ad
	PhotoLinks  string `json:"photoLinks"` // photoLinks element of ad
	Cost        uint   `json:"cost"`       // cost element of ad
}

// Stringers for ad struct
func (ad Ad) String() string {
	return fmt.Sprintf("Name: %s\nDescription: %s\nPhotoLinks: %s\nCost %d\n", ad.Name, ad.Description, ad.PhotoLinks, ad.Cost)
}

//get ad by id
func GetAd(id int) Ad {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") // open db
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var ad Ad // get ad
	err = db.QueryRow("SELECT `id`, `name`, `description`, `photoLinks`, `cost` FROM `announcementlist` WHERE `id` = ?", id).Scan(&ad.id, &ad.Name, &ad.Description, &ad.PhotoLinks, &ad.Cost)
	if err != nil {
		panic(err)
	}
	return ad
	//fmt.Printf("Name: %s\nDescription: %s\nPhotoLinks: %s\nCost %d\n\n", element.Name, element.Description, element.PhotoLinks, element.Cost)
}

//get all ads sorted by name/cost ASC(ascending)/DESC(descending)
func GetAllAds(field string, order string) []Ad {
	ads := []Ad{} //Create arr of ads

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") //open db
	if err != nil {
		panic(err)
	}

	defer db.Close()
	//get all ads
	res, err := db.Query(fmt.Sprintf("SELECT `name`, `description`, `photoLinks`, `cost` FROM `announcementlist` ORDER BY `%s` %s", field, order))
	if err != nil {
		panic(err)
	}
	// filling the ads array
	for res.Next() {
		var ad Ad
		err = res.Scan(&ad.Name, &ad.Description, &ad.PhotoLinks, &ad.Cost)
		if err != nil {
			panic(err)
		}
		ads = append(ads, ad)
		//fmt.Printf("Name: %s\nDescription: %s\nPhotoLinks: %s\nCost %d\n\n", list.Name, list.Description, list.PhotoLinks, list.Cost)
	}
	return ads
}

//(name, description, photoLinks, cost )
func CreateAd(name string, description string, photoLinks string, cost int) int64 {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") //open db
	if err != nil {
		panic(err)
	}

	defer db.Close()
	//inserting ad
	res, err := db.Exec(fmt.Sprintf("INSERT INTO `announcementlist` (`name`, `description`, `photoLinks`, `cost`) VALUES('%s','%s','%s','%d')", name, description, photoLinks, cost))
	if err != nil {
		panic(err)
	}

	//get id of inserting element
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return id
}

func main() {

	//fmt.Println(CreateAd("C", "asdas", "asd", 11))
	//fmt.Println(GetAd(1))
	// adlist := GetAllAds("name", "ASC") //(name/cost, ASC/DESC)
	// for _, value := range adlist {
	// 	fmt.Println(value)
	// }

}
