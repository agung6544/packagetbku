package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ayam struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Jenis          string             `bson:"jenis,omitempty" json:"jenis,omitempty"`
	Umur		   int             `bson:"umur,omitempty" json:"umur,omitempty"`
	Bobot          int             `bson:"bobot,omitempty" json:"bobot,omitempty"`
	Tinggi         int             `bson:"tinggi,omitempty" json:"tinggi,omitempty"`
	Jenis_Kelamin  string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
	Harga		   int             `bson:"harga,omitempty" json:"harga,omitempty"`
}

type Order struct {
	ID           	  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Ayam 	          Ayam               `bson:"ayam,omitempty" json:"ayam,omitempty"`
	Nama_Pemesan 	  string 			 `bson:"nama_pemesan,omitempty" json:"nama_pemesan,omitempty"`
	Alamat 			  string 			 `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tanggal_Pemesanan primitive.DateTime `bson:"tanggal_pemesanan,omitempty" json:"tanggal_pemesanan,omitempty"`
}