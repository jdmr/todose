package main

import "net/http"

func getHealthz(w http.ResponseWriter, r *http.Request) {
	err := client.Ping(r.Context(), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
