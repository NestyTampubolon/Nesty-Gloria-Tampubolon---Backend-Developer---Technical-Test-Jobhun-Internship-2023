package dto

type JurusanUpdateDTO struct {
	Id          uint64 `json:"id"`
	NamaJurusan string `json:"nama_jurusan" form:"nama_jurusan"`
}

type JurusanCreateDTO struct {
	NamaJurusan string `json:"nama_jurusan" form:"nama_jurusan" binding:"required" `
}
