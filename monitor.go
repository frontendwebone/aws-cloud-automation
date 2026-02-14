package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SystemHealth struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	Service   string `json:"service"`
}

func main() {
	// 1. Data Create Karein
	health := SystemHealth{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    "Healthy",
		Service:   "Cloud-Automation-Engine",
	}
	jsonData, _ := json.Marshal(health)

	// 2. Mock Cloud Server (Port 4566) par bhejein
	url := "http://localhost:4566/system-report.json"
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error: Is your Mock Server running on 4566?")
		return
	}
	defer resp.Body.Close()

	fmt.Println("Success! System Health Report uploaded to Cloud.")
}
