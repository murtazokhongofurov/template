package server

import (
	"database/sql"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/template/config"
	"github.com/template/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeOut     = 5
)

type Server struct {
	Cfg         *config.Config
	Db          *sql.DB
	RedisClient *redis.Client
	AwsClient   *minio.Client
	Logger      logger.Logger
}

func Run(*Server) *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
