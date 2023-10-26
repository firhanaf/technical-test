package data

import (
	"errors"
	"tech-test/features/mahasiswa"
	_nilaiData "tech-test/features/nilai"

	"gorm.io/gorm"
)

type MahasiswaQuery struct {
	db *gorm.DB
}

// GetById implements mahasiswa.MahasiswaDataInterface.
func (repo *MahasiswaQuery) GetById(NIM uint) (mahasiswa.CoreMahasiswa, error) {
	var data Mahasiswa
	tx := repo.db.Where("nim = ?", NIM).Preload("Nilais").Find(&data)
	if tx.Error != nil {
		return mahasiswa.CoreMahasiswa{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return mahasiswa.CoreMahasiswa{}, errors.New("data not found")
	}
	var mahasiswaData mahasiswa.CoreMahasiswa
	var nilaiMahasiswa []_nilaiData.CoreNilai

	for _, nilai := range data.Nilais { // Use data.Nilais instead of mahasiswaData.Nilais
		nilaiMahasiswa = append(nilaiMahasiswa, _nilaiData.CoreNilai{
			ID:             nilai.ID,
			UTS:            nilai.UTS,
			UAS:            nilai.UAS,
			AkhirAngka:     nilai.AkhirAngka,
			AkhirHuruf:     nilai.AkhirHuruf,
			MahasiswaNIM:   nilai.MahasiswaNIM,
			MataKuliahKode: nilai.MataKuliahKode,
		})
	}

	mahasiswaData = mahasiswa.CoreMahasiswa{ // Populate mahasiswaData here
		NIM:        data.NIM,
		Nama:       data.Nama,
		Kota:       data.Kota,
		TahunMasuk: data.TahunMasuk,
		Nilais:     nilaiMahasiswa,
	}

	return mahasiswaData, nil
}

func New(db *gorm.DB) mahasiswa.MahasiswaDataInterface {
	return &MahasiswaQuery{
		db: db,
	}
}
