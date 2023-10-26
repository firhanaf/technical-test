package handler

type NilaiResponse struct {
	ID             uint   `json:"id"`
	MataKuliahKode string `json:"kode_mata_kuliah"`
	AkhirHuruf     string `json:"nilai"`
}
