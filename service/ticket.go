package service

import (
	"GoBus/data"
	"GoBus/dto"
	"encoding/json"
	"errors"
	"strings"
)

type TicketService struct{}

// constructor
func NewTicketService() TicketService {
	return TicketService{}
}

// method buy ticket
func (ticketservice *TicketService) BuyTicket(req dto.TicketRequest) (dto.TicketResponse, error) {
	// validasi nama penumpang dan tujuan
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Destination) == "" {
		return dto.TicketResponse{}, errors.New("nama penumpang dan tujuan harus diisi")
	}

	// data harga tiket
	price := make(map[string]float64)
	err := json.Unmarshal([]byte(data.Destination), &price)
	if err != nil {
		return dto.TicketResponse{}, errors.New("failed: ambil data harga")
	}

	// cari harga
	var foundPrice float64
	var foundDestination string
	var isFound bool = false

	for destination, price := range price {
		if strings.EqualFold(destination, req.Destination) {
			foundPrice = price
			foundDestination = destination
			isFound = true
			break
		}
	}

	// jika tujuan tidak ditemukan
	if !isFound {
		return dto.TicketResponse{}, errors.New("tujuan tidak ditemukan")
	}

	// kirim response
	response := dto.TicketResponse{
		Name: req.Name,
		Destination: foundDestination,
		Price: foundPrice,
	}

	return  response, nil

}
