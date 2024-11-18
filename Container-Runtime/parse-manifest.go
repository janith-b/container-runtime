package main

import (
	"encoding/json"
	"fmt"
	"syscall"
)

type Platform struct {
	Architecture string `json:"architecture"`
	Os           string `json:"os"`
}

type Manifest struct {
	Digest    string   `json:"digest"`
	MediaType string   `json:"mediaType"`
	Platform  Platform `json:"platform"`
	Size      string   `json:"size"`
}

type Mainfests struct {
	Manifests []Manifest `json:"manifests"`
}

var manifests Mainfests

func parseManifest(bs []byte) {
	utsname := syscall.Utsname{}
	syscall.Uname(&utsname)
	arch_bs := []byte{}
	for _, e := range utsname.Machine {
		if e != 0 {
			arch_bs = append(arch_bs, byte(e))
		} else {
			continue
		}
	}
	json.Unmarshal(bs, &manifests)

	if string(arch_bs) == "x86_64" {
		for _, v := range manifests.Manifests {
			if v.Platform.Architecture == "amd64" && v.Platform.Os == "linux" {
				fmt.Println(v.Digest)
				break
			}
		}
	}
}
