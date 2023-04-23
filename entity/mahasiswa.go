package entity

import "time"

type Mahasiswa struct {
	Id                uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Nama              string    `gorm:"type:varchar(255)" json:"nama"`
	Usia              uint64    `gorm:"type:integer" json:"usia,string"`
	Gender            uint64    `gorm:"type:integer" json:"gender,string"`
	TanggalRegistrasi time.Time `gorm:"type:datetime" json:"tanggal_registrasi"`
}
