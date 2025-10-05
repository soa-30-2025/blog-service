package main

import (
    "log"
    "net"

    pb "blog-service/proto/blog"
    "blog-service/db"
    "blog-service/repository"
    "blog-service/services"
    "blog-service/handlers"

    "google.golang.org/grpc"
)

func main() {
    database := db.Connect()
    defer database.Close()

    repo := &repository.BlogRepository{DB: database}
    service := &services.BlogService{Repo: repo}
    handler := &handlers.BlogHandler{Service: service}

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Greška pri pokretanju servera: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterBlogServiceServer(grpcServer, handler)

    log.Println("Server pokrenut na :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Greška pri radu servera: %v", err)
    }
}
