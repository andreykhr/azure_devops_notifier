package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

type AzureDevOps struct {
	authorizeValue string
}

func (do *AzureDevOps) Init(login string, password string) {
	data := fmt.Sprintf("%s:%s", login, password)
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	do.authorizeValue = sEnc
}

func (do *AzureDevOps) GetWorkitems() {
	reader, err := os.Open("queries/query.json")

	request, err := http.NewRequest("POST", "https://do.norbit.ru/DIS/GPB/_apis/wit/wiql?api-version=6.0", reader)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Basic " + do.authorizeValue)

	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	dump, err := httputil.DumpRequest(request, true)

	fmt.Printf("%s\n",dump)

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Http status code wrong. Recived [%d], Expected [%d]", resp.StatusCode, http.StatusOK)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	log.Print(bodyString)
}