package entity

type Jurusan struct {
	Id          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NamaJurusan string `gorm:"type:varchar(255)" json:"nama_jurusan"`
}
