package main

import (
    "log"
    "net"

    pb "blog-service/proto/blog"
    "blog-service/repository"
    "blog-service/services"
    "blog-service/handlers"

    "google.golang.org/grpc"

	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"context"
)

func main() {

    dbUrl := os.Getenv("DATABASE_URL")
    if dbUrl == "" {
        dbUrl = "postgres://postgres:password@postgres:5432/blogdb?sslmode=disable"
    }

    pool, err := pgxpool.New(context.Background(), dbUrl)
    if err != nil {
        log.Fatalf("db connect: %v", err)
    }
    defer pool.Close()

    repo := &repository.BlogRepository{DB: pool}
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
