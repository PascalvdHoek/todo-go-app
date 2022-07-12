package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)

type Album struct {
	ID     int32   `json: "id"`
	Title  string  `json: "title"`
	Artist string  `json: "artist"`
	Price  float64 `json: "price"`
}

type City struct {
	ID          int32  `json: "id"`
	Name        string `json: "name"`
	CountryCode string `json: "countryCode"`
	District    string `json: "district"`
	Population  int32  `json: "population"`
}

type Country struct {
	Code           string   `json: "code"`
	Name           string   `json: "name"`
	Continent      string   `json: "continent"`
	Region         string   `json: "region"`
	SurfaceArea    float64  `json: "surfaceArea"`
	InDepYear      *int32   `json: "inDepYear"`
	Population     int32    `json: "population"`
	LifeExpectancy *float64 `json: "lifeExpectancy"`
	GNP            *float64 `json: "gnp"`
	GNPOld         *float64 `json: "gnpOld"`
	LocalName      string   `json: "localName"`
	GovernmentForm string   `json: "governmentForm"`
	HeadOfState    *string  `json: "headOfState"`
	Capital        *int32   `json: "capital"`
	Code2          string   `json: "code2"`
}

func main() {
	router := gin.Default()
	router.Routes()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.GET("/cities", getCities)
	router.GET("/countries", getCountries)
	//router.POST("/albums", postAlbum)
	//router.GET(("/albums/:id"), getAlbumById)
	router.Run("localhost:8080")

	fmt.Println(router.Routes())
}

func getCities(c *gin.Context) {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT id, name, countrycode, district, population FROM city ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var cities []City
	for rows.Next() {
		var city City
		rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		cities = append(cities, city)
	}
	c.IndentedJSON(http.StatusOK, cities)

}

func getAlbums(c *gin.Context) {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT id, title, artist, price FROM albums ")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var albums []Album
	for rows.Next() {
		var album Album
		err = rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		albums = append(albums, album)
	}
	c.IndentedJSON(http.StatusOK, albums)
}

func getCountries(c *gin.Context) {
	var (
		countries []Country
		db        *sql.DB
		rows      *sql.Rows

		err error
	)
	db, err = connectDB()
	if err != nil {
		fmt.Println(err)
	}
	rows, err = db.Query("Select code, name, continent, region, surfacearea, indepyear, population, lifeexpectancy, gnp, gnpold, localname, governmentform, headofstate, capital, code2  from country")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var country Country

		err = rows.Scan(&country.Code, &country.Name, &country.Continent, &country.Region, &country.SurfaceArea, &country.InDepYear, &country.Population, &country.LifeExpectancy, &country.GNP, &country.GNPOld, &country.LocalName, &country.GovernmentForm, &country.HeadOfState, &country.Capital, &country.Code2)
		if err != nil {
			panic(err)
		}
		countries = append(countries, country)
	}
	c.IndentedJSON(http.StatusOK, countries)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	var album Album

	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
	}
	row := db.QueryRow(fmt.Sprintf("SELECT id, title, artist, price FROM albums WHERE id =%s", id))
	err = row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err == sql.ErrNoRows {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album with id " + id + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

//func postAlbum(c *gin.Context) {
//	var newAlbum album
//	if err := c.BindJSON(&newAlbum); err != nil {
//		return
//	}
//	albums = append(albums, newAlbum)
//
//	c.IndentedJSON(http.StatusCreated, albums)
//}

func connectDB() (*sql.DB, error) {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "docker"
		dbname   = "world"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	if err != nil {
		return nil, err

	}

	fmt.Println("Succesfully connected!")
	return db, nil
}
