package main

import (
    "net/http"
    "database/sql"
    "fmt"

    "github.com/gin-gonic/gin"
    _ "github.com/lib/pq"
)

// Coiffeur represents data about a coiffeur.
type Coiffeur struct {
    ID     int64  `json:"id"`
    Name  string  `json:"name"`
    Location string  `json:"location"`
}

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", "jbriant", "dev")
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return db
}

// Function for handling messages
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

// getCoiffeurs responds with the list of all coiffeurs as JSON.
func getCoiffeurs(c *gin.Context) {
    db := setupDB()

    printMessage("Getting coiffeurs...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT osm_id AS id, name::text, st_astext(way) as location FROM planet_osm_point")

    // check errors
    checkErr(err)

    // var response []JsonResponse
    var coiffeurs []Coiffeur

    // Foreach movie
    for rows.Next() {
        var id int64
        var name string
        var location string

        err = rows.Scan(&id, &name, &location)

        // check errors
        checkErr(err)

        coiffeurs = append(coiffeurs, Coiffeur{ID: id, Name: name, Location: location})
    }

    c.IndentedJSON(http.StatusOK, coiffeurs)
}

func main() {
    router := gin.Default()
    router.GET("/coiffeurs", getCoiffeurs)
    router.Run("localhost:8080")
}


