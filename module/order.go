package module

import (
	"context"
	"errors"
	"fmt"
	"github.com/agung6544/packagetbku/model"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOrder(db *mongo.Database, col string, ayam model.Ayam, nama_pemesan string, alamat string) (insertedID primitive.ObjectID, err error) {
	orderku := bson.M{
		"ayam":    ayam,
		"nama_pemesan":     nama_pemesan,
		"alamat":    alamat,
		"tanggal_pemesanan": primitive.NewDateTimeFromTime(time.Now().UTC()),
	}
	result, err := db.Collection(col).InsertOne(context.Background(), orderku)
	if err != nil {
		fmt.Printf("InsertOrder: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllOrder(db *mongo.Database, col string) (data []model.Order) {
	karyawan := db.Collection(col)
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetOrderFromID(_id primitive.ObjectID, db *mongo.Database, col string) (order model.Order, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&order)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return order, fmt.Errorf("no data found for ID %s", _id)
		}
		return order, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return order, nil
}

func UpdateOrder(db *mongo.Database, col string, id primitive.ObjectID, ayam model.Ayam, nama_pemesan string, alamat string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
		"ayam":    ayam,
		"nama_pemesan":     nama_pemesan,
		"alamat":     alamat,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateOrder: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteOrderByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}