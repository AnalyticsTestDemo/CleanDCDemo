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

func ListWeatherData(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: In ListWeatherData..")
	ListofWeatherData:=GetWeatherData()	
    fmt.Println("Endpoint Hit: Data obtained from GetWeatherData.....now converting to json.")
	json.NewEncoder(w).Encode(ListofWeatherData)
}



func GetWeatherData()[]WeatherData {
	var ListofWeatherData []WeatherData
	var eachrow WeatherData

    fmt.Println("Endpoint Hit: in GetWeatherData ")

    db := dbConn()
    fmt.Println("Endpoint Hit: In GetWeatherData again")

    selDB, err := db.Query(" SELECT weatherdate,location,Temperature FROM Weatherdata")
    if err != nil {
        panic(err.Error())
    }
    defer selDB.Close()

    fmt.Println("Endpoint Hit: Query run on GetWeatherData")

    for selDB.Next() {    
        err = selDB.Scan( &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp)
        if err != nil {
 			log.Fatal(err)
        }         
        ListofWeatherData = append(ListofWeatherData, eachrow)
    }
    return ListofWeatherData
}




func AddWeatherData(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: In AddNewWeatherData..")
	InsertNewData()
    fmt.Println("Endpoint Hit:After insert ..")
}
func InsertNewData() {
	var newData WeatherData
    
    db := dbConn()
    fmt.Println("In insertweatherdata after opening connection")
  	
    newData.WeatherID = 1
    newData.WeatherDate = "06/06/2021"
    newData.Location= "Atlanta" 
    newData.Temp = "90"

    fmt.Println("In insertweatherdata before Prepare")

    insForm, err := db.Prepare(" INSERT INTO WeatherData (WeatherID,WeatherDate, Location,Temperature) VALUES (?,?,?,?)")
    if err != nil {
         log.Fatal(err)
         fmt.Println("ERror in inserting weather data")
    }

    insForm.Exec(newData.WeatherID, newData.WeatherDate,newData.Location,newData.Temp)
    log.Println("INSERT new WeatherData: Date: " + newData.WeatherDate + " | Loc: " + newData.Location)
	 
    defer db.Close()
}
