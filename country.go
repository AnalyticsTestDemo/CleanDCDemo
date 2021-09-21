package main

import (
    "fmt"
    "log"
    
    "net/http"
	//"database/sql"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
	"encoding/json"
)
func InsertNewCountryData(w http.ResponseWriter, r *http.Request) {
    var newData CountryMaster
    newData.CountryId = 3
    newData.CountryName="USA"

	fmt.Println("Endpoint Hit: In Insert new country..")
	db := dbConn()
    fmt.Println("Endpoint Hit: again in Insert new country after open connection..")

    
    fmt.Println("Endpoint Hit: In Insert new country.. before prepare")

    insForm, err := db.Prepare(" INSERT INTO CountryMaster(CountryID, CountryName) VALUES (?,?)")
    if err != nil {
        log.Fatal(err)
        fmt.Println("ERror in inserting new country")
    }
    insForm.Exec( newData.CountryId,newData.CountryName)
    fmt.Println("Inserted data for USA ")
    defer db.Close()
}

 
 func CountryList(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: In CountryList.. after running docker compose up")

	ListofCountries:=GetCountryData()	
    fmt.Println("Endpoint Hit: Data obtained from GetCountryData.....now converting to json.. docker compose up ")
	json.NewEncoder(w).Encode(ListofCountries)
}


func GetCountryData()[]CountryMaster {
	var ListofCountries []CountryMaster
	var eachrow CountryMaster

    fmt.Println("Endpoint Hit: in GetCountryData ")

    db := dbConn()
    fmt.Println("Endpoint Hit: In GetCountryData again")

    selDB, err := db.Query(" SELECT countryid, countryname FROM CountryMaster")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Endpoint Hit: Query run on Countrymaster")

    for selDB.Next() {    
        err = selDB.Scan(&eachrow.CountryId, &eachrow.CountryName )
        if err != nil {
 			log.Fatal(err)
        }         
        ListofCountries = append(ListofCountries, eachrow)
    }
    return ListofCountries
}
