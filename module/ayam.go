package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/agung6544/packagetbku/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertAyam(db *mongo.Database, col string, jenis string, umur int, bobot int, tinggi int, jenis_kelamin string, harga int) (insertedID primitive.ObjectID, err error) {
	ayamku := bson.M{
		"jenis":    jenis,
		"umur":     umur,
		"bobot":     bobot,
		"tinggi": tinggi,
		"jenis_kelamin":  jenis_kelamin,
		"harga":      harga,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), ayamku)
	if err != nil {
		fmt.Printf("InsertAyam: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllAyam(db *mongo.Database, col string) (data []model.Ayam) {
	ayamku := db.Collection(col)
	filter := bson.M{}
	cursor, err := ayamku.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAyamFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ayam model.Ayam, errs error) {
	ayamku := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := ayamku.FindOne(context.TODO(), filter).Decode(&ayam)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ayam, fmt.Errorf("no data found for ID %s", _id)
		}
		return ayam, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ayam, nil
}

func UpdateAyam(db *mongo.Database, col string, id primitive.ObjectID, jenis string, umur int, bobot int, tinggi int, jenis_kelamin string, harga int) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"jenis":    jenis,
			"umur":     umur,
			"bobot":    bobot,
			"tinggi":   tinggi,
			"jenis_kelamin":  jenis_kelamin,
			"harga":    harga,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateAyam: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteAyamByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	ayamku := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := ayamku.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}