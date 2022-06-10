package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/drewfrost/grpc-user-management/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	newUsers := make(map[string]int32)
	newUsers["Alice"] = 34
	newUsers["Bob"] = 32
	for name, age := range newUsers {
		user, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		log.Printf("Created user: %s with ID: %d", user.Name, user.Id)
	}

}
