package repository

import (
	"mahasiswa/entity"
	"gorm.io/gorm"
)

type JurusanRepository interface {
	InsertJurusan(b entity.Jurusan) entity.Jurusan
	UpdateJurusan(b entity.Jurusan) entity.Jurusan
	DeleteJurusan(b entity.Jurusan)
	AllJurusan() []entity.Jurusan
	FindJurusanID(JurusanID uint64) entity.Jurusan
}

type jurusanConnection struct {
	connection *gorm.DB
}

func NewJurusanRepository(dbConn *gorm.DB) JurusanRepository {
	return &jurusanConnection{
		connection: dbConn,
	}
}

func (db *jurusanConnection) InsertJurusan(b entity.Jurusan) entity.Jurusan {
	db.connection.Save(&b)
	return b
}

func (db *jurusanConnection) UpdateJurusan(b entity.Jurusan) entity.Jurusan {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *jurusanConnection) DeleteJurusan(b entity.Jurusan){
	db.connection.Delete(&b)
}

func (db *jurusanConnection) FindJurusanID(jurusanID uint64) entity.Jurusan {
	var jurusan entity.Jurusan
	db.connection.Find(&jurusan, jurusanID)
	return jurusan
}

func (db *jurusanConnection) AllJurusan() []entity.Jurusan {
	var jurusans []entity.Jurusan
	db.connection.Preload("User").Find(&jurusans)
	return jurusans
}
