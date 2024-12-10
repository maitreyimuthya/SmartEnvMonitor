package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type SensorData struct {
	Temperature float64 `json:"temperature"`
	Light       float64 `json:"light"`
	Humidity    float64 `json:"humidity"`
	Timestamp   string  `json:"timestamp"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/iotdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/sensor-data", handleSensorData)
	http.HandleFunc("/get-data", getData)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSensorData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data SensorData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	data.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO sensor_data (temperature, light, humidity, timestamp) VALUES (?, ?, ?, ?)",
		data.Temperature, data.Light, data.Humidity, data.Timestamp)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT temperature, light, humidity, timestamp FROM sensor_data ORDER BY timestamp DESC LIMIT 10")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []SensorData
	for rows.Next() {
		var sd SensorData
		if err := rows.Scan(&sd.Temperature, &sd.Light, &sd.Humidity, &sd.Timestamp); err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			return
		}
		data = append(data, sd)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
