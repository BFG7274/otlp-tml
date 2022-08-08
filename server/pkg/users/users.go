package users

import (
	context "context"
	"log"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserClient struct {
	userClient UserDataClient
}

func (uc *UserClient) GetUser(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := uc.userClient.GetUser(ctx, &GetUserRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *UserClient) GetUsers() (*UserList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := uc.userClient.GetUsers(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (uc *UserClient) UpdateUserMoney(u *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := uc.userClient.UpdateUserMoney(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func NewUserClient() *UserClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := NewUserDataClient(conn)
	return &UserClient{
		userClient: c,
	}
}
