package util

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
)

func Inspect(certName string) {
	cert, err := tls.LoadX509KeyPair(
		fmt.Sprintf("%s.crt", certName),
		fmt.Sprintf("%s.key", certName),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Access the certificate's fields
	for _, c := range cert.Certificate {
		leaf, err := x509.ParseCertificate(c)
		if err != nil {
			log.Fatal(err)
		}

		// print the whole leaf as pretty json
		fmt.Printf("%+v", asJSON(leaf))
	}
}

func asJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
