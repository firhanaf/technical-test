package data

type Nilai struct {
	ID             uint `gorm:"primaryKey"`
	UTS            uint
	UAS            uint
	AkhirAngka     uint
	AkhirHuruf     string
	MahasiswaNIM   uint
	MataKuliahKode string
}
