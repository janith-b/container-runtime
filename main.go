// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// type Platform struct {
// 	Architecture string `json:"architecture"`
// 	Os           string `json:"os"`
// }

// type Manifest struct {
// 	Digest    string   `json:"digest"`
// 	MediaType string   `json:"mediaType"`
// 	Platform  Platform `json:"platform"`
// 	Size      string   `json:"size"`
// }

// type Manifests struct {
// 	Manifests     []Manifest `json:"manifests"`
// 	MediaType     string     `json:"mediaType"`
// 	SchemaVersion string     `json:"schemaVersion"`
// }

// func main() {
// 	manifests := Manifests{}
// 	f, _ := os.ReadFile("redis-manifest.json")
// 	json.Unmarshal(f, &manifests)

// 	for _, v := range manifests.Manifests {
// 		fmt.Println(v.Platform.Os, v.Platform.Architecture)
// 	}

// 	fmt.Println(manifests.Manifests[0].Platform.Architecture)
// }
