package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting SSH Script Agent...")
	serveHTTP()
}

func serveHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/cmd/{command}", handleCommand).Methods("GET")

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", GetConfig().Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
	}()
	log.Println("HTTP Server listening")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	httpServer.Shutdown(ctx)
}

func isValidCommand(cmd string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
	return re.MatchString(cmd)
}

func isValidAuth(r *http.Request) bool {
	if GetConfig().Username == "" {
		return true
	}

	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return (user == GetConfig().Username) && (pass == GetConfig().Password)
}

func handleCommand(w http.ResponseWriter, r *http.Request) {
	if !isValidAuth(r) {
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	command := vars["command"]

	if !isValidCommand(command) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmd := exec.Command(GetConfig().Command, command)
	res, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
