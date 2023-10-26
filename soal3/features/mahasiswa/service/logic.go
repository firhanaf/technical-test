package service

import "tech-test/features/mahasiswa"

type MahasiswaService struct {
	mahasiswaData mahasiswa.MahasiswaDataInterface
}

// ReadById implements mahasiswa.MahasiswaServiceInterface.
func (service *MahasiswaService) ReadById(NIM uint) (mahasiswa.CoreMahasiswa, error) {
	result, err := service.mahasiswaData.GetById(NIM)
	return result, err
}

func New(repo mahasiswa.MahasiswaDataInterface) mahasiswa.MahasiswaServiceInterface {
	return &MahasiswaService{
		mahasiswaData: repo,
	}
}
