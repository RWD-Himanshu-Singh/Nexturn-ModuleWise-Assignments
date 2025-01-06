package main

import (
	"climate/services"
	"fmt"
)

func main() {
	// Display menu options
	for {
		fmt.Println("\nClimate Data Analysis")
		fmt.Println("1. Find City with Highest Temperature")
		fmt.Println("2. Find City with Lowest Temperature")
		fmt.Println("3. Calculate Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall")
		fmt.Println("5. Search City by Name")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			highest := services.FindHighestTemperature()
			fmt.Printf("City with highest temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)

		case 2:
			lowest := services.FindLowestTemperature()
			fmt.Printf("City with lowest temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)

		case 3:
			averageRainfall := services.CalculateAverageRainfall()
			fmt.Printf("Average Rainfall: %.2f mm\n", averageRainfall)

		case 4:
			var threshold float64
			fmt.Print("Enter rainfall threshold (mm): ")
			fmt.Scanln(&threshold)
			filteredCities, err := services.FilterCitiesByRainfall(threshold)
			if err != nil {
				fmt.Println(err)
			} else if len(filteredCities) == 0 {
				fmt.Println("No cities found with rainfall above the threshold.")
			} else {
				fmt.Println("Cities with rainfall above the threshold:")
				for _, city := range filteredCities {
					fmt.Printf("- %s (%.2f mm)\n", city.Name, city.Rainfall)
				}
			}

		case 5:
			var name string
			fmt.Print("Enter city name to search: ")
			fmt.Scanln(&name)
			city, err := services.SearchCityByName(name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("City: %s, Temperature: %.2f°C, Rainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			}

		case 6:
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
