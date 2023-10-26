package data

type MataKuliah struct {
	Kode         string `gorm:"primaryKey"`
	Nama         string
	SKS          uint
	DosenNIP     uint
	MahasiswaNIM uint
}
