package handler

import (
	"net/http"
	"strconv"
	"tech-test/features/mahasiswa"
	_nilai "tech-test/features/nilai/handler"
	"tech-test/helpers"

	"github.com/labstack/echo/v4"
)

type MahasiswaHandler struct {
	mahasiswaService mahasiswa.MahasiswaServiceInterface
}

func New(service mahasiswa.MahasiswaServiceInterface) *MahasiswaHandler {
	return &MahasiswaHandler{
		mahasiswaService: service,
	}
}

func (handler *MahasiswaHandler) GetMahasiswaData(c echo.Context) error {
	id := c.Param("nim")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error NIM not valid", nil))
	}
	result, err := handler.mahasiswaService.ReadById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error read mahasiswa data", nil))
	}
	var nilais []_nilai.NilaiResponse
	for _, nilai := range result.Nilais {
		nilais = append(nilais, _nilai.NilaiResponse{
			ID:             nilai.ID,
			AkhirHuruf:     nilai.AkhirHuruf,
			MataKuliahKode: nilai.MataKuliahKode,
		})
	}

	resultResponse := MahasiswaResponse{
		NIM:   result.NIM,
		Nama:  result.Nama,
		Nilai: nilais,
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read mahasiswa data", resultResponse))
}
