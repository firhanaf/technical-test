package data

import (
	"tech-test/features/mahasiswa"
	"tech-test/features/nilai"
	_dataNilai "tech-test/features/nilai/data"
)

type Mahasiswa struct {
	NIM        uint `gorm:"primaryKey;autoIncrement"`
	Nama       string
	Kota       string
	TahunMasuk string
	Nilais     []_dataNilai.Nilai `gorm:"foreignKey:MahasiswaNIM;references:NIM"`
}

func MahasiswaModelToCore(dataModel Mahasiswa) mahasiswa.CoreMahasiswa {
	return mahasiswa.CoreMahasiswa{
		NIM:        dataModel.NIM,
		Nama:       dataModel.Nama,
		Kota:       dataModel.Kota,
		TahunMasuk: dataModel.TahunMasuk,
		Nilais:     []nilai.CoreNilai{},
	}
}

func MahasiswaCoreToModel(dataCore mahasiswa.CoreMahasiswa) Mahasiswa {
	return Mahasiswa{
		NIM:        dataCore.NIM,
		Nama:       dataCore.Nama,
		Kota:       dataCore.Kota,
		TahunMasuk: dataCore.TahunMasuk,
		Nilais:     []_dataNilai.Nilai{},
	}
}
