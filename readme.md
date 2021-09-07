# TODO

1. Implementing swagger API Documentation

# Tasks accomplished

1. Implemention of 2 microservices communitating by event driven kafka messaging system(segmentio)
2. Dockerisation of application and script to rin docker file

# run application 

1. cd to directory having parking.sh
2. run 

    sudo ./parking.sh up

# service1(producer) structure

├── app
│   └── app.go
├── client
│   └── kafka
│       ├── config.go
│       └── kafka.go
├── config
│   ├── config.go
│   └── conf.json
├── controller
│   └── controller.go
├── docker-compose.yml
├── Dockerfile
├── docs
│   └── doc.go
├── go.mod
├── go.sum
├── main.go
├── models
│   └── models.go
└── service
    └── service.go

## Highlights

1. gin is used as web framework
2. Hexagonal way of implimentation is followed
3. code is highly modular and capable of relpacing and client implementation
4. viper is used for all the configration setup
5. uber/zap is used for logging
6. mongoDb is used as database
7. mongo Express is also provided for live feed of mongo data (localhost:8081)
8. kafca is used for inter service communication


# service2(consumer) structure

├── app
│   └── app.go
├── client
│   ├── kafka
│   │   └── kafka.go
│   └── mongo.go
├── config
│   ├── config.go
│   └── conf.json
├── controller
│   ├── controller.go
│   └── serviceController.go
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── repo
│   │   └── parking.go
│   └── req-res.go
└── service
    └── service.go

## Highlights

1. fin is used for web framework
2. mvc pattern is used for data flow


# curl request

## GetAllParkingSpots

curl --location --request GET 'localhost:9090/GetAllParkingSpots'

## GetEmptyParkingSpots

curl --location --request GET 'localhost:9090/GetEmptyParkingSpots'

## FillSlot

1. select any slot positions from [GetEmptyParkingSpots] and add to spotId field and also add vehicleId

curl --location --request POST 'localhost:8080/fillSlot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "spotId" : "s1",
    "vehicleId" : "v1"
}'

## EmptySlot

1. select and slots which are filled(can be identified using [GetAllParkingSpots] ) and populate spotId field of the request

curl --location --request POST 'localhost:8080/emptySlot' \
--header 'Content-Type: application/json' \
--data-raw '{
    "spotId" : "s1"
}'