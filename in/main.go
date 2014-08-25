package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/cheggaaa/pb"
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

	tarCmd := exec.Command("tar", "-C", destination, "-xf", "-")
	tarCmd.Stderr = os.Stderr

	tarIn, err := tarCmd.StdinPipe()
	if err != nil {
		fatal("creating tar pipe", err)
	}

	err = tarCmd.Start()
	if err != nil {
		fatal("starting tar", err)
	}

	resp, err := http.Get(request.Source.URI)
	if err != nil {
		fatal("requesting uri", err)
	}

	bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES)

	bar.Start()

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		fatal("unzipping", err)
	}

	_, err = io.Copy(io.MultiWriter(tarIn, bar), gz)
	if err != nil {
		fatal("downloading", err)
	}

	err = tarIn.Close()
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
