Create Database WeatherDB; 

CREATE TABLE CountryMaster
(
    CountryID bigint ,
    CountryName varchar(100)    
);

INSERT INTO CountryMaster(CountryID, CountryName) VALUES
 (1, 'Japan'),
 (2, 'China');


use WeatherDB 

CREATE TABLE StateMaster
(
    StateID bigint ,
    CountryID bigint,
    StateName varchar(100)    
);

INSERT INTO StateMaster(StateID, CountryId, StateName) VALUES(1,1, 'Saitama');
INSERT INTO StateMaster(StateID, CountryId, StateName) VALUES (2,1, 'Chiba');
INSERT INTO StateMaster(StateID, CountryId, StateName) VALUES (3,1, 'Kanagawa');
INSERT INTO StateMaster(StateID, CountryId, StateName) VALUES (4,2, 'Beijing');
INSERT INTO StateMaster(StateID, CountryId, StateName) VALUES (5,2, 'Shanghai');



CREATE TABLE LocationMaster
(
    LocationID bigint ,
    StateID bigint,
    LocationName varchar(100),
    Latitude FLOAT,
    Longitude Float    
);
INSERT INTO LocationMaster(LocationID,StateID, LocationName, Latitude,Longitude) 
VALUES(1,1,'Tokyo',35.6762,139.6503);

INSERT INTO LocationMaster(LocationID,StateID, LocationName, Latitude,Longitude) 
VALUES(2,1,'Kyoto',35.0116,135.7681);

INSERT INTO LocationMaster(LocationID,StateID, LocationName, Latitude,Longitude) 
VALUES(1,1,'Shenzhen',22.542883,114.062996);



Create Table WeatherData (
      WeatherID bigint,
      WeatherDate    varchar(100),
      Location varchar(300),
      Temperature varchar (300)

  )
--Insert into WeatherData(WeatherID ,WeatherDate  ,Location ,Temperature) values (1,'01/01/2021','Tokyo','107')
