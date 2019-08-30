package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	port := "8888"
	server := http.NewServeMux()

	server.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {
		setCors(&w)

		log.Printf("/play")

		data := struct {
			PlayerToken string `json:"playerToken"`
		}{
			PlayerToken: "eirik123",
		}

		dataJson, _ := json.Marshal(data)

		_, _ = w.Write(dataJson)
	})

	server.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		setCors(&w)

		log.Printf("/state")

		data := struct {
			Status string `json:"status"`
		}{
			Status: "",
		}

		dataJson, _ := json.Marshal(data)

		_, _ = w.Write(dataJson)
	})

	server.HandleFunc("/action", func(w http.ResponseWriter, r *http.Request) {
		setCors(&w)

		bytes, err := httputil.DumpRequest(r, true)
		if err != nil {
			panic(err)
		}
		log.Println(string(bytes))

		if err := r.ParseForm(); err != nil {
			panic(err)
		}

		keyPressed := r.Form.Get("keyPressed")
		playerToken := r.Form.Get("playerToken")

		log.Printf("/action keyPressed=%s playerToken=%s", keyPressed, playerToken)

		data := struct {
			Status      string `json:"status"`
			PlayerToken string `json:"playerToken"`
			KeyPressed  string `json:"keyPressed"`
		}{
			Status:      "playing",
			PlayerToken: playerToken,
			KeyPressed:  keyPressed,
		}

		dataJson, _ := json.Marshal(data)

		_, _ = w.Write(dataJson)
	})

	server.HandleFunc("/screen", func(w http.ResponseWriter, r *http.Request) {
		setCors(&w)
		_, _ = w.Write([]byte("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011100000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"))
	})

	log.Printf("mock server started on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}

func setCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
