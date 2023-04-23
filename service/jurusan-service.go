package service

import (
	"mahasiswa/dto"
	"mahasiswa/entity"
	"mahasiswa/repository"
	"log"
	"github.com/mashingan/smapping"
)

type JurusanService interface {
	Insert(b dto.JurusanCreateDTO) entity.Jurusan
	Update(b dto.JurusanUpdateDTO) entity.Jurusan
	Delete(b entity.Jurusan)
	All() []entity.Jurusan
	FIndById(jurusanID uint64) entity.Jurusan
}

type jurusanService struct {
	jurusanRepository repository.JurusanRepository
}

func NewJurusanService(jurusanRepo repository.JurusanRepository) JurusanService {
	return &jurusanService{
		jurusanRepository: jurusanRepo,
	}
}

func (service *jurusanService) Insert(b dto.JurusanCreateDTO) entity.Jurusan {
	jurusan := entity.Jurusan{}
	err := smapping.FillStruct(&jurusan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.jurusanRepository.InsertJurusan(jurusan)
	return res
}

func (service *jurusanService) Update(b dto.JurusanUpdateDTO) entity.Jurusan {
	jurusan := entity.Jurusan{}
	err := smapping.FillStruct(&jurusan, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.jurusanRepository.UpdateJurusan(jurusan)
	return res
}

func (service *jurusanService) Delete(b entity.Jurusan) {
	service.jurusanRepository.DeleteJurusan(b)
}

func (service *jurusanService) All() []entity.Jurusan {
	return service.jurusanRepository.AllJurusan()
}

func (service *jurusanService) FIndById(jurusanID uint64) entity.Jurusan {
	return service.jurusanRepository.FindJurusanID(jurusanID)
}
