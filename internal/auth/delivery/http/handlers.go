package http

import (
	"github.com/template/config"
	"github.com/template/internal/auth"
	"github.com/template/internal/session"
	"github.com/template/pkg/logger"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
	sessUC session.UCSession
	logger logger.Logger
}

// func NewAuthHandler(cfg *config.Config, authUC auth.UseCase, sessUC session.UCSession, log logger.Logger) auth.Handlers {
// 	return &authHandlers{cfg: cfg, authUC: authUC, sessUC: sessUC, logger: log}
// }

// func (h *authHandlers) Register() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request)  {
// 		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(w, r), "", )
// 	}
// }
