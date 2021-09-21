package main

import (
	//"database/sql"
	"fmt"
	//"log"
	"database/sql"
	"text/template"

	"net/http"
	//"encoding/json"
	//	"text/template"
   // "time"
   //"strconv" // New import
   // _ "github.com/denisenkom/go-mssqldb"
  
  )


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside Index....")
	var tmpl = template.Must(template.ParseGlob("*.tmpl"))

	var ListWeatherData []WeatherData
	var eachrow WeatherData
	filterID := r.URL.Query().Get("date")
	//log.Println("Filter Id = " , filterID)
    fmt.Println("Inside Index")

	if len(filterID)<=0 {
		db := dbConn()
		selDB, err := db.Query("select  weatherdate,location,Temperature from weatherdata ")
		if err != nil {
			panic(err.Error())
			//fmt.Println("Error in gettign weather data " + err.Error())
		}
		defer selDB.Close()

    	for selDB.Next() {
       
			err = selDB.Scan( &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
			if err != nil {
				fmt.Println("Error in reading weather data " + err.Error())
				//log.Fatal(err)
			}
			fmt.Println("After gettign data...")
        	ListWeatherData = append(ListWeatherData, eachrow)
    	}
    	tmpl.ExecuteTemplate(w, "Index",ListWeatherData)
		defer db.Close()
	}else{

		FilterRecords(filterID,w)
	}

   
}




func FilterRecords(wDate string,w http.ResponseWriter) {
	fmt.Println("Inside Filter....")
	var tmpl = template.Must(template.ParseGlob("*.tmpl"))

	var ListWeatherData []WeatherData
	var eachrow WeatherData
	fmt.Println("In Filter Show")
	db := dbConn()
	  
		
	selDB, err := db.Query("select  Location, WeatherDate, temperature from weatherdata where weatherdate =?", wDate)
	if err != nil {
		fmt.Println("ERror")
	}
	fmt.Println("In Filter Show before loop")
	for selDB.Next() {
		err = selDB.Scan( &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
		if err != nil {
			fmt.Println("ERror")
		}
		ListWeatherData = append(ListWeatherData, eachrow)        
	}
	tmpl.ExecuteTemplate(w, "FilterWeatherData", ListWeatherData)
	defer db.Close()
}
	
	
	
	func New(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Inside New....")
		var tmpl = template.Must(template.ParseGlob("*.tmpl"))
	
		tmpl.ExecuteTemplate(w, "New", nil)
	}
	
	func  IsRecordDuplicate(newID int,db *sql.DB) bool{	
		/*var count int
		var newDate="06/11/2021"
		var newLoc="Ireland"
		row := db.QueryRow("SELECT COUNT(*) FROM weatherdata where weatherdate=? and location=?", newDate,newLoc)
		err := row.Scan(&count)
		if err != nil {
			fmt.Println("error")
		}*/
		return false
	}
	 
	func InsertNewWeatherRecord(w http.ResponseWriter, r *http.Request) {
		fmt.Println("In InsertNewWeatherRecord:" )

		db := dbConn()
		isduplicate := IsRecordDuplicate(1,db)
		fmt.Println("In InsertNewWeatherRecord .. after open con" )

		if isduplicate{
			fmt.Println("Duplicate Record Exists for Ireland!")
 		}else{
		if r.Method == "POST" {		
			var newData WeatherData
			fmt.Println("Inside Post" )

			//var wID int64
			wDate := r.FormValue("txtwDate")
			loc := r.FormValue("txtLocation")
			Temperatures := r.FormValue("txtTemp")
			fmt.Println("INSERT: Date: " + wDate + " | Loc: " + loc + " temps: " + Temperatures)
	
			newData.WeatherID = 100
			newData.WeatherDate = wDate
			newData.Location= loc 
			newData.Temp = Temperatures
	
			insForm, err := db.Prepare(" INSERT INTO WeatherData (WeatherID,weatherdate, location,temperature) VALUES (?,?,?,?)")
			if err != nil {
			   // panic(err.Error())
				fmt.Println("error")	
			}
			fmt.Println("After Prepare: ")

			insForm.Exec(newData.WeatherID, newData.WeatherDate,newData.Location,newData.Temp)
			fmt.Println("INSERT: Date: " + newData.WeatherDate + " | Loc: " + newData.Location)
		}
			
		  
		}
		defer db.Close()
		//http.Redirect(w, r, "/", 301)
	}
	