package entity

type MahasiswaHobi struct {
	IdMahasiswa uint64 `gorm:"type:integer" json:"id_mahasiswa"`
	IdHobi      uint64 `gorm:"type:integer" json:"id_hobi"`

	Mahasiswa Mahasiswa `gorm:"foreignkey:id_mahasiswa;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"mahasiswa"`
	Hobi      Hobi      `gorm:"foreignkey:id_hobi;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"hobi"`
}
