package handler

import (
	"tech-test/features/mahasiswa"
	_nilai "tech-test/features/nilai/handler"
)

type MahasiswaResponse struct {
	NIM   uint                   `json:"nim"`
	Nama  string                 `json:"nama"`
	Nilai []_nilai.NilaiResponse `json:"hasil_studi"`
}

func MahasiswaModelToResponse(m mahasiswa.CoreMahasiswa) MahasiswaResponse {
	return MahasiswaResponse{
		NIM:   m.NIM,
		Nama:  m.Nama,
		Nilai: []_nilai.NilaiResponse{},
	}
}
