package database

import (
	"fmt"
	"tech-test/app/config"
	_mahasiswaData "tech-test/features/mahasiswa/data"
	_nilaiData "tech-test/features/nilai/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_mahasiswaData.Mahasiswa{}, &_nilaiData.Nilai{})

}
