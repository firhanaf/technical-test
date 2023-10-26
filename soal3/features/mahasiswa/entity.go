package mahasiswa

import (
	"tech-test/features/nilai"
)

type CoreMahasiswa struct {
	NIM        uint
	Nama       string
	Kota       string
	TahunMasuk string
	Nilais     []nilai.CoreNilai
}

type MahasiswaDataInterface interface {
	GetById(NIM uint) (CoreMahasiswa, error)
}

type MahasiswaServiceInterface interface {
	ReadById(NIM uint) (CoreMahasiswa, error)
}
