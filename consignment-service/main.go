package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	go_micro_srv_consignment "shippy/consignment-service/proto/consignment"
)

const (
	PORT = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("faild listen %v", err)
	}
	log.Printf("listen on %v\n", PORT)

	server := grpc.NewServer()
	repo := Repository{}

	go_micro_srv_consignment.RegisterShippingServiceServer(server, &service{repo: repo})
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve")
	}
}

type IRepository interface {
	Create(consignment *go_micro_srv_consignment.Consignment) (*go_micro_srv_consignment.Consignment, error)
}

type Repository struct {
	consignments []*go_micro_srv_consignment.Consignment
}

func (repo *Repository) Create(consignment *go_micro_srv_consignment.Consignment) (*go_micro_srv_consignment.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	fmt.Println(repo.consignments)
	return consignment, nil
}

func (repo *Repository) GetAll() []*go_micro_srv_consignment.Consignment {
	return repo.consignments
}

type service struct {
	repo Repository
}

func (s *service) CreateConsignment(ctx context.Context, req *go_micro_srv_consignment.Consignment) (*go_micro_srv_consignment.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	resp := &go_micro_srv_consignment.Response{
		Created:     true,
		Consignment: consignment,
	}

	return resp, nil
}

func (s *service) GetConsignments(context.Context, *go_micro_srv_consignment.GetRequest) (*go_micro_srv_consignment.Response, error) {
	allConsignments := s.repo.GetAll()
	resp := &go_micro_srv_consignment.Response{Consignments: allConsignments}

	return resp, nil
}
