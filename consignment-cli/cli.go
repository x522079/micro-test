package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	go_micro_srv_consignment "shippy/consignment-service/proto/consignment"
	"time"
)

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error %v\n", err)
	}

	defer conn.Close()

	client := go_micro_srv_consignment.NewShippingServiceClient(conn)
	consignment, err := parseFile(DEFAULT_INFO_FILE)
	if err != nil {
		log.Fatalf("parse fail %v", err)
	}

	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("request fail %v", err)
	}

	fmt.Println(time.Now().String())
	resp, _ = client.GetConsignments(context.Background(), &go_micro_srv_consignment.GetRequest{})

	fmt.Println(time.Now().String())
	for _, one := range resp.Consignments {
		log.Printf("%v\n", one)
	}
}

const (
	ADDRESS           = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment-cli/consignment.json"
)

func parseFile(fileName string) (*go_micro_srv_consignment.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	consignment := &go_micro_srv_consignment.Consignment{}
	err = json.Unmarshal(data, consignment)
	if err != nil {
		return nil, err
	}

	return consignment, nil
}
