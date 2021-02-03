package main

import (
	"os"

	"github.com/ElladanTasartir/codepix-go/application/grpc"
	"github.com/ElladanTasartir/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
