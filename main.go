package main

import (
    "log"
    "net"

    pb "blog-service/proto/blog"
    "blog-service/server"
    "blog-service/db"
    "google.golang.org/grpc"
)

func main() {
    database := db.Connect()
    defer database.Close()

    blogServer := &server.BlogServer{DB: database}

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Greška pri pokretanju servera: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterBlogServiceServer(grpcServer, blogServer)

    log.Println("Server pokrenut na :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Greška pri radu servera: %v", err)
    }
}
