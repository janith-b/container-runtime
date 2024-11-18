package main

import (
	"os"
)

var args []string = os.Args

func main() {
	switch args[1] {
	case "pull":
		// pullImage("https://registry-1.docker.io/v2/", args[2], args[3])
		_, repo, tag := unmarshallImageName(args[2])
		pullManifest("https://registry-1.docker.io/v2/", repo, tag)
	}
}

func run() {

}
