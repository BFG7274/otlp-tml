package main

import (
	pb "bfg7274/otlp-tml-store/pkg/store"
	"context"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = 50052
)

type server struct {
	pb.UnimplementedStoreDataServer
}

type Data struct {
	ItemMap map[int64]pb.Item
}

func (d *Data) addItem(item *pb.Item) error {
	item.Id = int64(idNext)
	idNext++
	if _, ok := d.ItemMap[item.Id]; ok {
		return fmt.Errorf("user already exists")
	}
	d.ItemMap[item.Id] = *item
	return nil
}

func (d *Data) getItem(id int64) (*pb.Item, error) {
	if user, ok := d.ItemMap[id]; ok {
		return &user, nil
	}
	return nil, fmt.Errorf("user does not exist")
}

func (d *Data) getItems() ([]*pb.Item, error) {
	var r []*pb.Item
	for _, v := range d.ItemMap {
		r = append(r, &v)
	}
	return r, nil
}

func (d *Data) updateItemNum(item *pb.Item) error {
	if i, ok := d.ItemMap[item.Id]; ok {
		i.Num = item.Num
		return nil
	}
	return fmt.Errorf("user does not exist")
}

var idNext int
var data Data

func (s *server) GetItems(ctx context.Context, e *emptypb.Empty) (*pb.ItemList, error) {
	items, _ := data.getItems()
	return &pb.ItemList{
		Item: items,
	}, nil
}

func (s *server) GetItem(ctx context.Context, getItemRequest *pb.GetItemRequest) (*pb.Item, error) {
	item, err := data.getItem(getItemRequest.GetId())
	if err != nil {
		return &pb.Item{}, nil
	}
	return item, nil

}

func (s *server) UpdateItem(ctx context.Context, item *pb.Item) (*pb.StoreResponse, error) {
	err := data.updateItemNum(item)
	if err != nil {
		return &pb.StoreResponse{
			Status: false,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.StoreResponse{
		Status: true,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, item *pb.Item) (*pb.StoreResponse, error) {
	return &pb.StoreResponse{}, nil
}

func main() {
	data = Data{
		ItemMap: map[int64]pb.Item{},
	}
	data.addItem(&pb.Item{
		Name:  "Water",
		Price: 10,
	})
	data.addItem(&pb.Item{
		Name:  "Candy",
		Price: 20,
	})
	data.addItem(&pb.Item{
		Name:  "Breed",
		Price: 25,
	})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStoreDataServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
