/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/review-service/domain"
	pb "github.com/review-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReviewServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	GetBookRequest(ctx, c, 1)
	GetUserRequest(ctx, c, 1)

}

func GetBookRequest(ctx context.Context, conn pb.ReviewServiceClient, id uint) *domain.BookInfo {

	book, err := conn.GetBook(ctx, &pb.GetBookRequest{Id: int32(id)})

	if err != nil {
		log.Fatalf("could not receive: %v", err)
	}
	log.Printf("Received title: %s", book.GetTitle())
	log.Printf("Received author: %s", book.GetAuthor())

	return &domain.BookInfo{
		Title:  book.GetTitle(),
		Author: book.GetAuthor(),
	}
}

func GetUserRequest(ctx context.Context, conn pb.ReviewServiceClient, id uint) *domain.UserInfo {

	user, err := conn.GetUser(ctx, &pb.GetUserRequest{Id: int32(id)})

	if err != nil {
		log.Fatalf("could not receive: %v", err)
	}
	log.Printf("Received name: %s", user.GetName())

	return &domain.UserInfo{
		Name: user.GetName(),
	}
}
