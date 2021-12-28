package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zawachte/caas-rest-api/pkg/caas"
)

func main() {

	/**
		sampleaccount := caas.Account{
			Email:    "zach@zach.com",
			Username: "zawachte1",
			Password: "saltsalt",
		}

		b, err := json.Marshal(sampleaccount)
		if err != nil {
			log.Fatal(err)
		}

		resp, err := http.Post("http://localhost:8080/v1/account", "application/json", strings.NewReader(string(b)))
		if err != nil {
			log.Fatal(err)
		}

		log.Println(resp)


	sampleCluster := caas.Cluster{
		AccountId:  "0002",
		Kubeconfig: "sdfsdfsdfsdfd",
	}

	b, err := json.Marshal(sampleCluster)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/v1/cluster", "application/json", strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)

	**/

	resp, err := http.Get("http://localhost:8080/v1/account")
	if err != nil {
		log.Fatal(err)
	}

	a := []caas.Account{}
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(a, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
	log.Println(string(b))

	resp, err = http.Get("http://localhost:8080/v1/cluster")
	if err != nil {
		log.Fatal(err)
	}

	d := []caas.Cluster{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(d, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
	log.Println(string(b))

	resp, err = http.Get("http://localhost:8080/v1/cluster/findByAccountId?accountid=0002")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)

	d = []caas.Cluster{}
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		log.Fatal(err)
	}

	b, err = json.MarshalIndent(d, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
	log.Println(string(b))

}
