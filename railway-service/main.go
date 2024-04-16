package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/rs/cors"
)

var Train1 = Train{
	ID:          1,
	Name:        "Kandy Express",
	Departure:   "Colombo",
	Destination: "Kandy",
	Capacity:    100,
	Type:        "Express",
	Price:       1000,
}

var Train2 = Train{
	ID:          2,
	Name:        "Galle Express",
	Departure:   "Colombo",
	Destination: "Galle",
	Capacity:    100,
	Type:        "Express",
	Price:       1500,
}

var Train3 = Train{
	ID:          3,
	Name:        "Jaffna Railways",
	Departure:   "Colombo",
	Destination: "Jaffna",
	Capacity:    100,
	Type:        "Non-Express",
	Price:       800,
}

func main() {
	trains := addTrains()
	reservations := []Reservation{}

	mux := http.NewServeMux()
	mux.HandleFunc("/trains", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(trains)
		}
	})

	mux.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			var booking Booking
			err = json.Unmarshal(body, &booking)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			reservation, intError := addBooking(booking)
			if intError != nil {
				http.Error(w, intError.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(reservation)
			reservations = append(reservations, reservation)
		}
	})

	mux.HandleFunc("/reservations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(reservations)
		}
	})

	corsHandler := cors.AllowAll()
	handler := corsHandler.Handler(mux)

	fmt.Println("Server is starting on port 8080...")
	http.ListenAndServe(":8080", handler)
}

func addTrains() []Train {
	trains := []Train{}
	trains = append(trains, Train1, Train2, Train3)

	return trains
}

func addBooking(booking Booking) (Reservation, error) {
	var reservation Reservation
	uuid, _ := uuid.NewV7()
	if booking.Train == "Kandy Express" {
		reservation = Reservation{
			ID:     uuid.String(),
			Train:  Train1,
			Amount: booking.Amount,
			Cost:   booking.Amount * 1000,
		}
	} else if booking.Train == "Galle Express" {
		reservation = Reservation{
			ID:     uuid.String(),
			Train:  Train2,
			Amount: booking.Amount,
			Cost:   booking.Amount * 1500,
		}
	} else if booking.Train == "Jaffna Railways" {
		reservation = Reservation{
			ID:     uuid.String(),
			Train:  Train3,
			Amount: booking.Amount,
			Cost:   booking.Amount * 800,
		}
	} else {
		return Reservation{}, errors.New("Train not found")
	}

	return reservation, nil
}
