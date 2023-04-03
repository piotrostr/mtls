package client

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Run() {
	cert, err := os.ReadFile("client.crt")
	if err != nil {
		fmt.Println("Error loading client certificate")
		return
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	config := &tls.Config{
		RootCAs: certPool,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		fmt.Println("Error making request", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return
	}

	fmt.Println(string(body))
}
