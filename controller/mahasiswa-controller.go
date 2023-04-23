package controller

import (
	"github.com/gin-gonic/gin"
	"mahasiswa/dto"
	"mahasiswa/entity"
	"mahasiswa/helper"
	"mahasiswa/service"
	"net/http"
	"strconv"
)

type MahasiswaController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	UpdateStok(context *gin.Context)
	Delete(context *gin.Context)
}

type mahasiswaController struct {
	mahasiswaService service.MahasiswaService
}

func NewMahasiswaController(mahasiswaServ service.MahasiswaService) MahasiswaController {
	return &mahasiswaController{
		mahasiswaService: mahasiswaServ,
	}
}

func (c *mahasiswaController) All(context *gin.Context) {
	var mahasiswas []entity.Mahasiswa = c.mahasiswaService.All()
	// res := helper.BuildResponse(true, "OK", produks)
	context.JSON(http.StatusOK, mahasiswas)
}

func (c *mahasiswaController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var mahasiswa entity.Mahasiswa = c.mahasiswaService.FIndById(id)
	if (mahasiswa == entity.Mahasiswa{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		// res := helper.BuildResponse(true, "OK", produk)
		context.JSON(http.StatusOK, mahasiswa)
	}
}

func (c *mahasiswaController) Insert(context *gin.Context) {
	var mahasiswaCreateDTO dto.MahasiswaCreateDTO
	errDTO := context.ShouldBind(&mahasiswaCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.mahasiswaService.Insert(mahasiswaCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *mahasiswaController) Update(context *gin.Context) {
	var mahasiswaUpdateDTO dto.MahasiswaUpdateDTO
	errDTO := context.ShouldBind(&mahasiswaUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	mahasiswaUpdateDTO.Id = id
	result := c.mahasiswaService.Update(mahasiswaUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}

func (c *mahasiswaController) Delete(context *gin.Context) {
	var mahasiswa entity.Mahasiswa
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	mahasiswa.Id = id
	c.mahasiswaService.Delete(mahasiswa)
	res := helper.BuildResponse(true, "Delete", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)

}

func (c *mahasiswaController) UpdateStok(context *gin.Context) {
	var mahasiswaUpdateDTO dto.MahasiswaUpdateDTO
	errDTO := context.ShouldBind(&mahasiswaUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}

	var mahasiswa entity.Mahasiswa = c.mahasiswaService.FIndById(id)
	if (mahasiswa == entity.Mahasiswa{}) {
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	}
	mahasiswaUpdateDTO.Id = id
	mahasiswaUpdateDTO.Nama = mahasiswa.Nama
	mahasiswaUpdateDTO.Usia = mahasiswa.Usia
	mahasiswaUpdateDTO.Gender = mahasiswa.Gender
	mahasiswaUpdateDTO.TanggalRegistrasi = mahasiswa.TanggalRegistrasi
	result := c.mahasiswaService.Update(mahasiswaUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)

}
