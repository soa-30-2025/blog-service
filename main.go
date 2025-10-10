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
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatalln("Error while loading .env file")
    }

	dbUrl := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("Neuspešno povezivanje sa bazom: %v", err)
	}
	defer pool.Close()

	blogRepo := &repository.BlogRepository{DB: pool}
	commentRepo := &repository.CommentRepository{DB: pool}
    likeRepo := &repository.LikeRepository{DB: pool}

	blogService := &services.BlogService{Repo: blogRepo}
	commentService := &services.CommentService{Repo: commentRepo}
    likeService := &services.LikeService{Repo: likeRepo}

	handler := &handlers.BlogHandler{
		BlogService:    blogService,
		CommentService: commentService,
        LikeService: likeService,
	}
    
	port:= os.Getenv("PORT")
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalf("Greška pri pokretanju servera: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, handler)

	log.Println("gRPC server pokrenut na portu " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Greška pri radu servera: %v", err)
	}
}
