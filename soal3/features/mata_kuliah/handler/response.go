package handler

type MataKuliahResponse struct {
	Kode     uint   `json:"kode"`
	Nama     string `json:"mata_kuliah"`
	DosenNIP uint   `json:"nip"`
	SKS      uint   `json:"sks"`
}
