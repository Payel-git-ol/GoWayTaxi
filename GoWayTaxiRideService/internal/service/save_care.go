package service

import (
	"RideService/pkg/database"
	"RideService/pkg/models"
)

func SaveCar() {
	var CarMakes = []string{
		"Toyota Corolla", "Toyota Camry", "Hyundai Elantra", "Hyundai Sonata",
		"Kia Rio", "Kia Optima", "Volkswagen Polo", "Skoda Octavia",
		"Ford Focus", "Nissan Altima", "Honda Civic", "Renault Logan",
	}

	var CarColors = []string{
		"Yellow", "White", "Black", "Silver", "Gray", "Blue", "Green", "Red",
	}

	for i := 0; i < len(CarMakes) && i < len(CarColors); i++ {
		car := &models.Car{
			CarMake: CarMakes[i],
			Color:   CarColors[i],
		}
		database.DB.Create(car)
	}
}
