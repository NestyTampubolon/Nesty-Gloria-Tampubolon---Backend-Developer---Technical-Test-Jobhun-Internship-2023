package entity

type Hobi struct {
	Id       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	NamaHobi string `gorm:"type:varchar(255)" json:"nama_hobi"`
}
