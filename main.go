package main

import (
	"mahasiswa/config"
	"mahasiswa/controller"
	"mahasiswa/repository"
	"mahasiswa/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConnection()
	mahasiswaRepository repository.MahasiswaRepository = repository.NewMahasiswaRepository(db)
	mahasiswaService    service.MahasiswaService       = service.NewMahasiswaService(mahasiswaRepository)
	mahasiswaController controller.MahasiswaController = controller.NewMahasiswaController(mahasiswaService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	mahasiswaRoutes := r.Group("api/mahasiswas")
	{
		mahasiswaRoutes.GET("/", mahasiswaController.All)
		mahasiswaRoutes.POST("/", mahasiswaController.Insert)
		mahasiswaRoutes.GET("/:id", mahasiswaController.FindByID)
		mahasiswaRoutes.PUT("/:id", mahasiswaController.Update)
		mahasiswaRoutes.DELETE("/:id", mahasiswaController.Delete)
	}

	r.Run()
}
