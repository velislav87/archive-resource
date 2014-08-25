package main

import (
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

	wgetPipe := exec.Command(
		"sh",
		"-c",
		"wget -O - $1 | gunzip | tar -C $2 -xf -",
		"sh", request.Source.URI, destination,
	)
	wgetPipe.Stdout = os.Stderr
	wgetPipe.Stderr = os.Stderr

	err = wgetPipe.Run()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	json.NewEncoder(os.Stdout).Encode(models.InResponse{})
}

func fatal(doing string, err error) {
	println("error " + doing + ": " + err.Error())
	os.Exit(1)
}
