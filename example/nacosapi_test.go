package example

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNacosApiCreate(t*testing.T) {
	url := "http://47.115.24.72:8848/nacos/v1/ns/instance?ip=127.0.0.1&port=3026&enable=true&healthy=true&serviceName=jmfen-sport-scoreapi&ephemeral=false"
	method := "POST"

	payload := strings.NewReader("")

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}