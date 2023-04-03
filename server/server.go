package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func Run() {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println("Error loading server certificate and key")
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.TLS.PeerCertificates[0].Subject.CommonName)
	})

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: config,
	}

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Error starting server")
		return
	}
}
