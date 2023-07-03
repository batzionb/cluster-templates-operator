package main

import (
	"fmt"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"time"

	helmserver "github.com/stolostron/cluster-templates-operator/testutils/helm"
)

func main() {
	go forever()
	var server *httptest.Server
	var httpsServer *httptest.Server
	server = helmserver.StartGitRepoServer()
	httpsServer = helmserver.StartTLSGitRepoServer()
	fmt.Println(server.URL)
	fmt.Println(httpsServer.URL)
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Adios!")
}

func forever() {
	for {
		fmt.Printf("%v+\n", time.Now())
		time.Sleep(time.Second)
	}
}
