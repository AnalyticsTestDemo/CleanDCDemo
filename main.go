package main

import (
    "fmt"
    "log"
    "time"
    "os"
    "net/http"
	//"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	//"encoding/json"
	"database/sql"
)

type CountryMaster struct {
    CountryId    int
    CountryName  string
  }
 
type  Location struct {
    Cityname string
    State string
    Latitude float64
    Longitude float64
}

type WeatherData struct {
    WeatherID int
    Location string 
    WeatherDate string 
    Temp string
}

func main() {
	fmt.Println("Hi There! Welcome to Docker Demo. Time is : "+ time.Now().String())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 	fmt.Fprintf(w,"Hello world. How are you? " + time.Now().String())
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi There. How is it going? " + time.Now().String())
    })
    handleRequests()
}
 
 
func handleRequests() {
 
    http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/CountryList", CountryList)
    http.HandleFunc("/NewCountry",InsertNewCountryData)

    http.HandleFunc("/WeatherList", ListWeatherData)
    http.HandleFunc("/AddNew",AddWeatherData)

    http.HandleFunc("/WeatherUI",Index)
    http.HandleFunc("/WeatherUINew",New)
    http.HandleFunc("/WeatherUIAddNew",InsertNewWeatherRecord)


    
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request){
     fmt.Fprintf(w,"Hello.. This is my first app.. "+ time.Now().String())
 }




func dbConn() (db *sql.DB) {
    var condb *sql.DB
    db_user := os.Getenv("DB_USER")
    db_password := os.Getenv("DB_PASSWORD")

    fmt.Println("Endpoint Hit: In dbConn connection.. after runnign docker compose up")
    fmt.Println("Usenamd =" + db_user + ", pwd :" + db_password)

    var (
            server   string = "db"        // for example
            user     string = db_user     // Database user
            password string = db_password // User Password
            port     int    = 1433        // Database port
            database string = "WeatherDB"
    )

    connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
    condb, err := sql.Open("mssql", connString)

    // Test if the connection is OK or not
    if err != nil {
            panic("Cannot connect to database")
    } else {
            fmt.Println("Connected!")

    }

    if err = condb.Ping(); err != nil {
            condb.Close()
            fmt.Println("Endpoint Hit: Error in dbcon")
            fmt.Println(err)
    }

    // Make sure to update the Password value below from "Your_password123" to your actual password.
    // var connection = @"Server=db;Database=master;User=sa;Password=Your_password123;";

    fmt.Println("Endpoint Hit: After sql.Open")

    if err != nil {
            log.Fatal(err)
    }
    return condb
}


//WeatherDB is created under MASTER
//running http://localhost:8080/CountryList should give countrylist 
//docker run -w /go/src/app -it --link mysql55c -d --name golangapp -v $(pwd):/go/src/app golang bash -c "go get github.com/go-sql-driver/mysql;go build main.go; go test -v --config ./config.ini"
//SQL IP address             "IPAddress": "172.17.0.2",
//                        "HostPort": "1433"
// 842c6bfaea2c  -cleandemo 
// 56beb1db7406  - sql 2910 

/*
-----------
Tony's working code 
func dbConn() (db *sql.DB) {
        var condb *sql.DB
        db_user := os.Getenv("DB_USER")
        db_password := os.Getenv("DB_PASSWORD")

        fmt.Println("Endpoint Hit: In dbConn connection.. after runnign docker compose up")
        fmt.Println("Usenamd =" + db_user + ", pwd :" + db_password)

        var (
                server   string = "db"        // for example
                user     string = db_user     // Database user
                password string = db_password // User Password
                port     int    = 1433        // Database port
        )

        connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
        condb, err := sql.Open("mssql", connString)

        // Test if the connection is OK or not
        if err != nil {
                panic("Cannot connect to database")
        } else {
                fmt.Println("Connected!")

        }

        if err = condb.Ping(); err != nil {
                condb.Close()
                fmt.Println("Endpoint Hit: Error in dbcon")
                fmt.Println(err)
        }

        // Make sure to update the Password value below from "Your_password123" to your actual password.
        // var connection = @"Server=db;Database=master;User=sa;Password=Your_password123;";

        fmt.Println("Endpoint Hit: After sql.Open")

        if err != nil {
                log.Fatal(err)
        }
        return condb
}

*/