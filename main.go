package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Roll is model for sushi
type Roll struct {
	ID          string `json:"id"`
	ImageNumber string `json:"imageNumber"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
}

//Init rolls var as a slice
var rolls []Roll

//Get All Rolls
func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}

//Get Single Roll
func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//Create a New Roll
func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

//Update
func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

//Delete Roll
func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(rolls)
}

func main() {
	//stump data TODO: Implement database
	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "2", ImageNumber: "4", Name: "California Roll", Ingredients: "Crab, Cucumber, Avocado, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "3", ImageNumber: "10", Name: "Caterpillar Roll", Ingredients: "Tempura Shrimp, Cucumber, Avocado, Nori, Rice, Eel Sauce"})

	rolls = append(rolls, Roll{ID: "4", ImageNumber: "7", Name: "Spider Roll", Ingredients: "Tempura Crab, Nori, Rice, Avocado, Spicy mayonnaise"})

	rolls = append(rolls, Roll{ID: "5", ImageNumber: "2", Name: "Unagi Roll", Ingredients: "Broiled Eel, Avocado, Nori, Rice, Eel sauce"})

	rolls = append(rolls, Roll{ID: "6", ImageNumber: "1", Name: "Philidelphia Roll", Ingredients: "Smoked Salmon, Cream Cheese, Cucumber, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "7", ImageNumber: "3", Name: "Salmon Roll", Ingredients: "Salmon, Nori, Rice, Tobiko"})

	rolls = append(rolls, Roll{ID: "8", ImageNumber: "6", Name: "Avocado Roll", Ingredients: "Avocado, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "9", ImageNumber: "9", Name: "Rainbow Roll", Ingredients: "Yellow Tail, Salmon, Cucumber, Nori, Rice"})

	rolls = append(rolls, Roll{ID: "10", ImageNumber: "5", Name: "Tobiko Roll", Ingredients: "Tobiko, Nori, Sushi Rice"})

	rolls = append(rolls, Roll{ID: "11", ImageNumber: "11", Name: "Yellow Tail Sushi", Ingredients: "Yellow Tail, Wasabi, Sushi Rice"})
	//initialize router
	router := mux.NewRouter()

	//endpoints
	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	//added CORS for heroku deploy
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))
}
