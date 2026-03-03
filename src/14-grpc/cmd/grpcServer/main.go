package main

import (
	"database/sql"
	"net"

	"github.com/felipecaue-lb/pos-go-expert/src/14-grpc/internal/database"
	"github.com/felipecaue-lb/pos-go-expert/src/14-grpc/internal/pb"
	"github.com/felipecaue-lb/pos-go-expert/src/14-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
