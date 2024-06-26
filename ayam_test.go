package packagetbku

import (
	"fmt"
	"testing"

	"github.com/agung6544/packagetbku/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertAyam(t *testing.T) {
	jenis := "Mangon"
	umur := "8"
	bobot := "3"
	tinggi := "60"
	jenis_kelamin := "jantan"
	harga := "250000"
	insertedID, err := module.InsertAyam(module.MongoConn, "ayamku", jenis, umur, bobot, tinggi, jenis_kelamin, harga)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetAyamFromID(t *testing.T) {
	id := "667995e5b1769afb138d97b9"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetAyamFromID(objectID, module.MongoConn, "ayamku")
	if err != nil {
		t.Fatalf("error calling GetAyamFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAllAyam(t *testing.T) {
	data := module.GetAllAyam(module.MongoConn, "ayamku")
	fmt.Println(data)
}

func TestDeleteAyamByID(t *testing.T) {
	id := "667995e5b1769afb138d97b9" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteAyamByID(objectID, module.MongoConn, "ayamku")
	if err != nil {
		t.Fatalf("error calling DeleteAyamByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetAyamFromID
	_, err = module.GetAyamFromID(objectID, module.MongoConn, "ayamku")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}