package main

import (
	"compress/gzip"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/concourse/archive-resource/models"
)

func main() {
	if len(os.Args) < 2 {
		println("usage: " + os.Args[0] + " <destination>")
		os.Exit(1)
	}

	destination := os.Args[1]

	err := os.MkdirAll(destination, 0755)
	if err != nil {
		fatal("creating destination", err)
	}

	var request models.InRequest

	err = json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fatal("reading request", err)
	}

	// must use wget as net/http forces dynamic linking
	//
	// :(
	wgetCmd := exec.Command("wget", "-O", "-", request.Source.URI)
	wgetCmd.Stderr = os.Stderr

	wgetOut, err := wgetCmd.StdoutPipe()
	if err != nil {
		fatal("creating wget pipe", err)
	}

	gunzip, err := gzip.NewReader(wgetOut)
	if err != nil {
		fatal("creating gzip reader", err)
	}

	tarCmd := exec.Command("tar", "-C", destination, "-xf", "-")
	tarCmd.Stderr = os.Stderr
	tarCmd.Stdin = gunzip

	err = tarCmd.Start()
	if err != nil {
		fatal("starting tar", err)
	}

	err = wgetCmd.Run()
	if err != nil {
		fatal("closing tar stream", err)
	}

	err = tarCmd.Wait()
	if err != nil {
		fatal("untarring", err)
	}

	json.NewEncoder(os.Stdout).Encode(models.InResponse{})
}

func fatal(doing string, err error) {
	println("error " + doing + ": " + err.Error())
	os.Exit(1)
}
