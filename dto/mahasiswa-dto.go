package dto

import "time"

type MahasiswaUpdateDTO struct {
	Id                uint64    `json:"id"`
	Nama              string    `json:"nama" form:"nama"`
	Usia              uint64    `json:"usia,string" form:"usia`
	Gender            uint64    `json:"gender,string" form:"gender"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi,string" form:tanggal_registrasi"`
}

type MahasiswaCreateDTO struct {
	Nama              string    `json:"nama" form:"nama" binding:"required" `
	Usia              uint64    `json:"usia,string" form:"usia binding:"required"`
	Gender            uint64    `json:"gender,string" form:"gender" binding:"required"`
	TanggalRegistrasi time.Time `json:"tanggal_registrasi" form:tanggal_registrasi" binding:"required"`
}
