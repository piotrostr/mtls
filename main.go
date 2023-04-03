package main

import (
	"flag"

	"github.com/piotrostr/mtls/client"
	"github.com/piotrostr/mtls/server"
	"github.com/piotrostr/mtls/util"
)

var (
	runServer = flag.Bool("server", false, "Run server")
	inspect   = flag.Bool("inspect", false, "Inspect certificates")
	certName  = flag.String("cert", "client", "Certificate name")
)

func main() {
	flag.Parse()

	if *inspect && *certName != "" {
		util.Inspect(*certName)
		return
	}

	if *runServer {
		server.Run()
	} else {
		client.Run()
	}

}
