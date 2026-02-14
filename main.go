package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到GitLab CI Go服务演示!\n")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)
	
	port := "8082"
	log.Printf("服务器启动，监听端口 %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
