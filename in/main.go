package main

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
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

	sourceURL, err := url.Parse(request.Source.URI)
	if err != nil {
		fatal("parsing uri", err)
	}

	// busybox's wget does not support basic auth in the uri,
	// so convert it to the header format ourselves
	authHeader := "Authorization: "
	if sourceURL.User != nil {
		authHeader += "Basic " + basicAuth(sourceURL.User)
		sourceURL.User = nil
	}

	wgetPipe := exec.Command(
		"sh",
		"-c",
		"wget --header \"$3\" -O - \"$1\" | gunzip | tar -C \"$2\" -xf -",
		"sh", sourceURL.String(), destination, authHeader,
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

func basicAuth(user *url.Userinfo) string {
	username := user.Username()
	password, _ := user.Password()
	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}
