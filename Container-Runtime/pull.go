package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Auth struct {
	Token        string `json:"token"`
	Access_Token string `json:"access_token"`
	Expires_In   int    `json:"expires_in"`
	Issued_At    string `json:"issued_at"`
}

func unmarshallImageName(image string) (string, string, string) {
	//docker.io/library/redis:latest
	tag := strings.Split(image, ":")[1]
	imageRef := strings.Split(strings.Split(image, ":")[0], "/")
	registry := imageRef[0]
	namespace := imageRef[1]
	imageName := imageRef[2]
	repository := namespace + "/" + imageName
	return registry, repository, tag
}

func pull(url string, token string) ([]byte, int) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("An ERROR Occurred : ", err)
	}
	buff := strings.Builder{}
	_, err1 := io.Copy(&buff, resp.Body)
	if err1 != nil {
		fmt.Println("An ERROR Occurred : ", err1)
	}
	return []byte(buff.String()), resp.StatusCode
}

func pullManifest(registryEndpoint string, repo string, tag string) []byte {
	endpoint := registryEndpoint + repo + "/manifests/" + tag
	response, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("An ERROR Occurred : ", err)
	}
	if response.StatusCode == 401 {
		www_authenticate_header := strings.Split(strings.Split(strings.Replace(response.Header["Www-Authenticate"][0], "\"", "", -1), " ")[1], ",")
		authEndpoint := strings.Split(www_authenticate_header[0], "=")[1] + "?" + www_authenticate_header[1] + "&" + www_authenticate_header[2]
		token := getauthToken(authEndpoint)
		bs, _ := pull(endpoint, token)
		return bs
	} else {
		return make([]byte, 0)
	}
}

func getauthToken(authEndpoint string) string {
	//Generate Auth Endpoint
	response, err := http.Get(authEndpoint)
	if err != nil {
		fmt.Println("An ERROR Occurred : ", err)
	}

	buff := strings.Builder{}

	_, err1 := io.Copy(&buff, response.Body)
	if err != nil {
		fmt.Println("An ERROR Occurred : ", err1)
	}

	var authResponse Auth
	err2 := json.Unmarshal([]byte(buff.String()), &authResponse)
	if err != nil {
		fmt.Println("An ERROR Occurred : ", err2)
	}

	return authResponse.Token
}
