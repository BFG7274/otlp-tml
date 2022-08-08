package main

import (
	pb "bfg7274/otlp-tml-store/pkg/users"
	"context"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	port = 50051
)

type server struct {
	pb.UnimplementedUserDataServer
}

type User struct {
	id    int
	name  string
	money int
}

type Data struct {
	UserMap map[int]User
}

func (d *Data) addUser(u *User) error {
	u.id = idNext
	idNext++
	if _, ok := d.UserMap[u.id]; ok {
		return fmt.Errorf("user already exists")
	}
	d.UserMap[u.id] = *u
	return nil
}

func (d *Data) getUser(id int) (*User, error) {
	if user, ok := d.UserMap[id]; ok {
		return &user, nil
	}
	return nil, fmt.Errorf("user does not exist")
}

func (d *Data) getUsers() ([]User, error) {
	var r []User
	for _, v := range d.UserMap {
		r = append(r, v)
	}
	return r, nil
}

func (d *Data) updateUserMoney(u *User) error {
	if user, ok := d.UserMap[u.id]; ok {
		user.money = u.money
		return nil
	}
	return fmt.Errorf("user does not exist")
}

var idNext int
var data Data

func (s *server) GetUsers(ctx context.Context, e *emptypb.Empty) (*pb.UserList, error) {
	users, _ := data.getUsers()
	var u []*pb.User
	for _, v := range users {
		u = append(u, &pb.User{
			Id:    int64(v.id),
			Name:  v.name,
			Money: int64(v.money),
		})
	}
	return &pb.UserList{
		User: u,
	}, nil
}

func (s *server) GetUser(ctx context.Context, getUserRequest *pb.GetUserRequest) (*pb.User, error) {
	user, err := data.getUser(int(getUserRequest.GetId()))
	if err != nil {
		return &pb.User{}, nil
	}
	return &pb.User{
		Id:    int64(user.id),
		Name:  user.name,
		Money: int64(user.money),
	}, nil

}

func (s *server) UpdateUserMoney(ctx context.Context, user *pb.User) (*pb.Response, error) {
	err := data.updateUserMoney(&User{
		id:    int(user.GetId()),
		name:  user.GetName(),
		money: int(user.GetMoney()),
	})
	if err != nil {
		return &pb.Response{
			Status: false,
			Msg:    err.Error(),
		}, nil
	}
	return &pb.Response{
		Status: true,
	}, nil
}

func (s *server) DeleteUser(ctx context.Context, user *pb.User) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func main() {
	data = Data{
		UserMap: map[int]User{},
	}
	data.addUser(&User{
		name:  "Zhang San",
		money: 100,
	})
	data.addUser(&User{
		name:  "Li Si",
		money: 5000,
	})
	data.addUser(&User{
		name:  "Wang Wu",
		money: 30,
	})
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserDataServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
