package server

import (
	"database/sql"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/template/config"
	"github.com/template/pkg/logger"
)

type Server struct {
	Cfg         *config.Config
	Db          *sql.DB
	RedisClient *redis.Client
	AwsClient   *minio.Client
	Logger      logger.Logger
}

func Run(s *Server) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	
	return mux
}
