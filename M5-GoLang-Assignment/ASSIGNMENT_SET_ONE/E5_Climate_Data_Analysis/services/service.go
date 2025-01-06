package services

import (
	"climate/models"
	"errors"
)

var cities = []models.City{
	{Name: "Mumbai", Temperature: 20.5, Rainfall: 300},
	{Name: "Delhi", Temperature: 18.2, Rainfall: 250},
	{Name: "Chennai", Temperature: 19.8, Rainfall: 275},
	{Name: "Kolkata", Temperature: 21.2, Rainfall: 300},
	{Name: "Bhubaneswar", Temperature: 22.3, Rainfall: 280},
	{Name: "Noida", Temperature: 19.7, Rainfall: 250},
	{Name: "Itanagar", Temperature: 23.1, Rainfall: 285},
	{Name: "Puri", Temperature: 23.7, Rainfall: 290},
	{Name: "Balasore", Temperature: 21.6, Rainfall: 270},
}

// Find the city with the highest temperature
func FindHighestTemperature() models.City {
	highest := cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
	}
	return highest
}

// Find the city with the lowest temperature
func FindLowestTemperature() models.City {
	lowest := cities[0]
	for _, city := range cities {
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return lowest
}

// Calculate the average rainfall
func CalculateAverageRainfall() float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Filter cities by rainfall threshold
func FilterCitiesByRainfall(threshold float64) ([]models.City, error) {
	if threshold < 0 {
		return nil, errors.New("rainfall threshold cannot be negative")
	}
	var filteredCities []models.City
	for _, city := range cities {
		if city.Rainfall > threshold {
			filteredCities = append(filteredCities, city)
		}
	}
	return filteredCities, nil
}

// Search for a city by name
func SearchCityByName(name string) (models.City, error) {
	for _, city := range cities {
		if city.Name == name {
			return city, nil
		}
	}
	return models.City{}, errors.New("city not found")
}
