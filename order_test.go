package packagetbku

import (
	"fmt"
	"testing"

	"github.com/agung6544/packagetbku/model"
	"github.com/agung6544/packagetbku/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertOrder(t *testing.T) {
	ayam := model.Ayam{
		Jenis:     "Mangon",
		Umur: "8",
		Bobot:   "3",
		Tinggi:  "60",
		Jenis_Kelamin:  "jantan",
		Harga:  "250000",
	}
	nama_pemesan := "yudha"
	alamat := "8"
	insertedID, err := module.InsertOrder(module.MongoConn, "orderku", ayam, nama_pemesan, alamat)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetOrderFromID(t *testing.T) {
	id := "667c021ffd34f3142a49f087"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetOrderFromID(objectID, module.MongoConn, "orderku")
	if err != nil {
		t.Fatalf("error calling GetOrderFromID: %v", err)
	}
	fmt.Println(biodata)
}

func TestGetAllOrder(t *testing.T) {
	data := module.GetAllOrder(module.MongoConn, "orderku")
	fmt.Println(data)
}

func TestDeleteOrderByID(t *testing.T) {
	id := "667c021ffd34f3142a49f087" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteOrderByID(objectID, module.MongoConn, "orderku")
	if err != nil {
		t.Fatalf("error calling DeleteOrderByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetOrderFromID
	_, err = module.GetOrderFromID(objectID, module.MongoConn, "orderku")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}