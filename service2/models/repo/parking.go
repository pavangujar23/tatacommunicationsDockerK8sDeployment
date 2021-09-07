package repo

import (
	"context"
	"fmt"
	"strconv"

	"github.com/service2/client"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Parking Exported
type Parking struct {
	SpotId    string `json:"spotId" bson:"_id,omitempty"`
	VehicleId string `json:"vehicleId"`
	Occupied  string `json:"occupied"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

var parkingCollection *mongo.Collection = client.GetParkingCollection()

func init() {
	parkingCollection.Drop(context.Background())
	PostInitialParkingSpots()
}

// PostInitialParkingSpots saves initial paring spots to mongo
func PostInitialParkingSpots() {
	for i := 1; i <= 10; i++ {
		PostParkingSpot(Parking{"s" + strconv.Itoa(i), "", "false", 0, 0})
	}
}

// GetParkingByID takes id and returns parking struct
func GetParkingByID(id string) (parking *Parking, err error) {

	ctx := context.Background()

	cur := parkingCollection.FindOne(ctx, bson.D{{"_id", id}})
	cur.Decode(&parking)
	if parking == nil {
		return nil, fmt.Errorf("No parking object found")
	}
	return parking, nil
}

// PostParkingSpot takes parking struct and saves to mongo
func PostParkingSpot(parking Parking) error {
	ctx := context.Background()
	_, err := parkingCollection.InsertOne(ctx, parking)
	if err != nil {
		fmt.Println("Error while inserting data to Mongo")
		return err
	}
	return nil
}

// PutParkingSpot updates existing parking struct with new one
func PutParkingSpot(parking Parking) error {
	ctx := context.Background()
	filter := bson.D{{"_id", parking.SpotId}}
	update := bson.M{
		"$set": parking,
	}
	_, err := parkingCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("Error while updating parking data")
	}
	return nil
}

// GetAllParkingSpots returns list of all parking structs
func GetAllParkingSpots() (parkingList []Parking, err error) {

	ctx := context.Background()

	cursor, err := parkingCollection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Error while getting data from Mongo")
		return nil, fmt.Errorf("Error while getting data from Mongo")
	}
	if err := cursor.All(context.Background(), &parkingList); err != nil {
		fmt.Println("Error while converting cursor data to List of parking")
		return nil, fmt.Errorf("Error while converting cursor data to List of parking")
	}
	if parkingList == nil {
		fmt.Println("No parking spots found")
		return nil, fmt.Errorf("No parking spots found")
	}
	return parkingList, nil
}

// GetEmptyParkingSpots returns list of all empty parking structs
func GetEmptyParkingSpots() (parkingList []Parking, err error) {

	ctx := context.Background()

	cursor, err := parkingCollection.Find(ctx, bson.D{{"occupied", "false"}})
	if err != nil {
		fmt.Println("Error while getting data from Mongo")
		return nil, fmt.Errorf("Error while getting data from Mongo")
	}
	if err = cursor.All(context.Background(), &parkingList); err != nil {
		fmt.Println("Error while converting cursor data to List of parking")
		return nil, fmt.Errorf("Error while converting cursor data to List of parking")

	}
	if parkingList == nil {
		fmt.Println("No empty parking spots found")
		return nil, fmt.Errorf("No empty parking spots found")

	}
	return parkingList, nil
}
