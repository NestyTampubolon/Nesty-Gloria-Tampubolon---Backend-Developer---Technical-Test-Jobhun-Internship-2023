package repository

import (
	"mahasiswa/entity"
	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	InsertMahasiswa(b entity.Mahasiswa) entity.Mahasiswa
	UpdateMahasiswa(b entity.Mahasiswa) entity.Mahasiswa
	DeleteMahasiswa(b entity.Mahasiswa)
	AllMahasiswa() []entity.Mahasiswa
	FindMahasiswaID(MahasiswaID uint64) entity.Mahasiswa
}

type mahasiswaConnection struct {
	connection *gorm.DB
}

func NewMahasiswaRepository(dbConn *gorm.DB) MahasiswaRepository {
	return &mahasiswaConnection{
		connection: dbConn,
	}
}

func (db *mahasiswaConnection) InsertMahasiswa(b entity.Mahasiswa) entity.Mahasiswa {
	db.connection.Save(&b)
	return b
}

func (db *mahasiswaConnection) UpdateMahasiswa(b entity.Mahasiswa) entity.Mahasiswa {
	db.connection.Save(&b)
	db.connection.Find(&b)
	return b
}

func (db *mahasiswaConnection) DeleteMahasiswa(b entity.Mahasiswa){
	db.connection.Delete(&b)
}

func (db *mahasiswaConnection) FindMahasiswaID(mahasiswaID uint64) entity.Mahasiswa {
	var mahasiswa entity.Mahasiswa
	db.connection.Find(&mahasiswa, mahasiswaID)
	return mahasiswa
}

func (db *mahasiswaConnection) AllMahasiswa() []entity.Mahasiswa {
	var mahasiswas []entity.Mahasiswa
	db.connection.Preload("User").Find(&mahasiswas)
	return mahasiswas
}
