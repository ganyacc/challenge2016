package main

import (
	"log"
	"net/http"

	"movie-distribution/handlers"
	"movie-distribution/models"
)

func main() {
	// Load locations from CSV
	err := models.LoadLocationsFromCSV("cities.csv")
	if err != nil {
		log.Fatalf("Failed to load locations: %v", err)
	}

	// Initialize distributors
	d1 := models.NewDistributor("DISTRIBUTOR1", nil)
	d1.AddPermission("INCLUDE", "IN")        // India
	d1.AddPermission("INCLUDE", "US")        // United States
	d1.AddPermission("EXCLUDE", "KA-IN")     // Karnataka, India
	d1.AddPermission("EXCLUDE", "CHE-TN-IN") // Chennai, Tamil Nadu, India

	d2 := models.NewDistributor("DISTRIBUTOR2", d1)
	d2.AddPermission("INCLUDE", "IN")
	d2.AddPermission("EXCLUDE", "TN-IN")

	d3 := models.NewDistributor("DISTRIBUTOR3", d2)
	d3.AddPermission("INCLUDE", "HBL-KA-IN")

	// Initialize handler with distributors
	h := handlers.NewHandler(map[string]*models.Distributor{
		"DISTRIBUTOR1": d1,
		"DISTRIBUTOR2": d2,
		"DISTRIBUTOR3": d3,
	})

	// Set up routes
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/check-permission", h.CheckPermission)

	// Start server
	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
