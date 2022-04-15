# Drone Navigation Service API

## Story
Space exploration is underway and mostly done by private companies. You joined
Engineering department of one of the private government contractors, Atlas Corporation.
In that year and age, everything is automated, so survey and data gathering is done by drones.
Unfortunately, although drones are perfectly capable of gathering the data, they have issues
with locating databank to upload gathered data. You, as the most promising recruit of Atlas
Corp, were tasked to develop a drone navigation service (DNS):
- each observed sector of the galaxy has unique numeric SectorID assigned to it
- each sector will have at least one DNS deployed
- each sector has different number of drones deployed at any given moment
- it’s future, but not that far, so drones will still use JSON REST API

## Goal
To create a microservice API that can helps drones get location of databank.



Installation Requirements:
-----------------

1. go 1.17+ 

How to run:
-----------------

1. Clone the repo

	git clone https://github.com/kumareswaramoorthi/dns.git

2. Navigate to project directory 

	cd dns 

3. Build the application by following command

	go build -o dns main.go

4. Run the application by the following command 

	./dns 


Alternatively, using docker,


1. Clone the repo

	git clone https://github.com/kumareswaramoorthi/dns.git

2. Navigate to project directory 

	cd dns 

3. Build the docker image by following command

	docker build -t dns:1.0 .

4. Run the application by the following command 

	docker run -d -p 8080:8080 dns:1.0


## **Swagger**

Swagger UI can be accessed at http://127.0.0.1:8080/swagger/index.html


## Documentation for API Endpoints

All URIs are relative to *http://127.0.0.1:8080*


## **1.Get Location**

Method | HTTP request | Description
------------- | ------------- | -------------
**GetLocation** | **POST** /location | Get Location of Databank

### Parameters

JSON body containing the coordinates information

### Response 

 - HTTP Status 200.
 - JSON object containing location of the databank.


### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### Example request and response

 - **Request**: `curl -H "Content-type: application/json" -d '{"x": "123.12","y": "456.56","z": "789.89","vel": "20.0"}' 127.0.01:8080/location`
 - **Response**: `{"loc":1389.57}`


## **2.Health Check**

Method | HTTP request | Description
------------- | ------------- | -------------
**HealthCheck** | **GET** / | Health Check API


### Parameters
This endpoint does not need any parameter.

### Response
HTTP Status 200.




