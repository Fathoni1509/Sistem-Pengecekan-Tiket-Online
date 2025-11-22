package main

import (
	"GoBus/dto"
	"GoBus/handler"
	"GoBus/service"
)

func main() {
	// inisialisasi
	ticketService := service.NewTicketService()
	ticketHandler := handler.NewTicketHandler(ticketService)

	// buat request
	passenger := dto.TicketRequest {
		Name: "Habibi",
		Destination: "Jakarta",
	}

	// panggil method buy
	ticketHandler.Buy(passenger)
}