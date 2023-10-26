package data

type Dosen struct {
	NIP  uint `gorm:"primaryKey"`
	Nama string
	Kota string
}
