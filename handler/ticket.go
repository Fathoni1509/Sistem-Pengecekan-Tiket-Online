package handler

import (
	"GoBus/dto"
	"GoBus/service"
	"fmt"
)

type TicketHandler struct {
	Ticketservice service.TicketService
}

// constructor
func NewTicketHandler(ticketservice service.TicketService) TicketHandler {
	return  TicketHandler{
		Ticketservice: ticketservice,
	}
}

// method buy ticket
func (tickethandler *TicketHandler) Buy(req dto.TicketRequest) {
	response, err := tickethandler.Ticketservice.BuyTicket(req)

	if err != nil {
		fmt.Println("Gagal membeli tiket")
		return
	}

	// cetak hasil
	fmt.Println("\n=== Harga Tiket ===")
	fmt.Printf("Penumpang : %s\n", response.Name)
	fmt.Printf("Tujuan    : %s\n", response.Destination)
	fmt.Printf("Harga     : Rp %.2f\n", response.Price)
	fmt.Println("===================")
}