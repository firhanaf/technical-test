package router

import (
	"tech-test/app/config"
	_mahasiswaData "tech-test/features/mahasiswa/data"
	_mahasiswaHandler "tech-test/features/mahasiswa/handler"
	_mahasiswaService "tech-test/features/mahasiswa/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(cfg *config.AppConfig, c *echo.Echo, db *gorm.DB) {
	MahasiswaData := _mahasiswaData.New(db)
	MahasiswaService := _mahasiswaService.New(MahasiswaData)
	MahasiswaHandlerAPI := _mahasiswaHandler.New(MahasiswaService)

	c.GET("/soal3/datamahasiswa/:nim", MahasiswaHandlerAPI.GetMahasiswaData)
}
