package store

import (
	context "context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StoreClient struct {
	userClient StoreDataClient
}

func (sc *StoreClient) GetItem(id int) (*Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := sc.userClient.GetItem(ctx, &GetItemRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *StoreClient) GetUsers() (*ItemList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := uc.userClient.GetItems(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *StoreClient) UpdateUserMoney(item *Item) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := uc.userClient.UpdateItem(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func NewStoreClient() *StoreClient {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := NewStoreDataClient(conn)
	return &StoreClient{
		userClient: c,
	}
}
