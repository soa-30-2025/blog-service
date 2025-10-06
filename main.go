package main

import (
	"context"
	"log"
	"net"
	"os"

	"blog-service/handlers"
	pb "blog-service/proto/blog"
	"blog-service/repository"
	"blog-service/services"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:password@postgres:5432/blogdb?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("❌ Neuspešno povezivanje sa bazom: %v", err)
	}
	defer pool.Close()

	// Repozitorijumi
	blogRepo := &repository.BlogRepository{DB: pool}
	commentRepo := &repository.CommentRepository{DB: pool}
    likeRepo := &repository.LikeRepository{DB: pool}

	// Servisi
	blogService := &services.BlogService{Repo: blogRepo}
	commentService := &services.CommentService{Repo: commentRepo}
    likeService := &services.LikeService{Repo: likeRepo}

	// Handler koji sadrži oba servisa
	handler := &handlers.BlogHandler{
		BlogService:    blogService,
		CommentService: commentService,
        LikeService: likeService,
	}

	// Pokretanje gRPC servera
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Greška pri pokretanju servera: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, handler)

	log.Println("✅ gRPC server pokrenut na portu :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Greška pri radu servera: %v", err)
	}
}
