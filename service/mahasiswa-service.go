package service

import (
	"mahasiswa/dto"
	"mahasiswa/entity"
	"mahasiswa/repository"
	"log"
	"github.com/mashingan/smapping"
)

type MahasiswaService interface {
	Insert(b dto.MahasiswaCreateDTO) entity.Mahasiswa
	Update(b dto.MahasiswaUpdateDTO) entity.Mahasiswa
	Delete(b entity.Mahasiswa)
	All() []entity.Mahasiswa
	FIndById(mahasiswaID uint64) entity.Mahasiswa
}

type mahasiswaService struct {
	mahasiswaRepository repository.MahasiswaRepository
}

func NewMahasiswaService(mahasiswaRepo repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{
		mahasiswaRepository: mahasiswaRepo,
	}
}

func (service *mahasiswaService) Insert(b dto.MahasiswaCreateDTO) entity.Mahasiswa {
	mahasiswa := entity.Mahasiswa{}
	err := smapping.FillStruct(&mahasiswa, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.mahasiswaRepository.InsertMahasiswa(mahasiswa)
	return res
}

func (service *mahasiswaService) Update(b dto.MahasiswaUpdateDTO) entity.Mahasiswa {
	mahasiswa := entity.Mahasiswa{}
	err := smapping.FillStruct(&mahasiswa, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.mahasiswaRepository.UpdateMahasiswa(mahasiswa)
	return res
}

func (service *mahasiswaService) Delete(b entity.Mahasiswa) {
	service.mahasiswaRepository.DeleteMahasiswa(b)
}

func (service *mahasiswaService) All() []entity.Mahasiswa {
	return service.mahasiswaRepository.AllMahasiswa()
}

func (service *mahasiswaService) FIndById(mahasiswaID uint64) entity.Mahasiswa {
	return service.mahasiswaRepository.FindMahasiswaID(mahasiswaID)
}
