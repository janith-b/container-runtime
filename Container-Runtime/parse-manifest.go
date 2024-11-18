package main

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
