package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo"
	httperrors "github.com/template/pkg/httpErrors"

	"github.com/pkg/errors"
	"github.com/template/config"
	"github.com/template/internal/models"
	"github.com/template/pkg/logger"
	"github.com/template/pkg/sanitize"
)

const HeaderXRequestID = "X-Request-ID"

// Get request id from echo context
func GetRequestID(w http.ResponseWriter, r *http.Request) string {
	return r.Header.Get(HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from echo context
func GetCtxWithReqID(w http.ResponseWriter, r *http.Request) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(r.Response.Request.Context(), time.Second*15)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(w, r))
	return ctx, cancel
}

// Get context  with request id
func GetRequestCtx(w http.ResponseWriter, r *http.Request) context.Context {

	return context.WithValue(r.Context(), ReqIDCtxKey{}, GetRequestID(w, r))
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

// Configure jwt cookie
func ConfigureJWTCookie(cfg *config.Config, jwtToken string) *http.Cookie {
	return &http.Cookie{
		Name:       cfg.Cookie.Name,
		Value:      jwtToken,
		Path:       "/",
		RawExpires: "",
		MaxAge:     cfg.Cookie.MaxAge,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HttpOnly,
		SameSite:   0,
	}
}

// Configure jwt cookie
func CreateSessionCookie(cfg *config.Config, session string) *http.Cookie {
	return &http.Cookie{
		Name:  cfg.Session.Name,
		Value: session,
		Path:  "/",
		// Domain: "/",
		// Expires:    time.Now().Add(1 * time.Minute),
		RawExpires: "",
		MaxAge:     cfg.Session.Expire,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HttpOnly,
		SameSite:   0,
	}
}

// Delete session
func DeleteSessionCookie(w http.ResponseWriter, sessionName string) {
	cookie := &http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

// UserCtxKey is a key used for the User object in the context
type UserCtxKey struct{}

// Get user from context
func GetUserFromCtx(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(UserCtxKey{}).(*models.User)
	if !ok {
		return nil, httperrors.Unauthorized
	}

	return user, nil
}

// Get user ip address
func GetIPAddress(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

// Error response with logging error for echo context
func ErrResponseWithLog(w http.ResponseWriter, r *http.Request, logger logger.Logger, err error) error {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(w, r),
		GetIPAddress(r),
		err,
	)
	// return ctx.JSON(httpErrors.ErrorResponse(err))
	
	return w.Write(httperrors.ErrorResponse(err))
}

// Error response with logging error for echo context
func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// Read request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

func ReadImage(ctx echo.Context, field string) (*multipart.FileHeader, error) {
	image, err := ctx.FormFile(field)
	if err != nil {
		return nil, errors.WithMessage(err, "ctx.FormFile")
	}

	// Check content type of image
	if err = CheckImageContentType(image); err != nil {
		return nil, err
	}

	return image, nil
}

// Read sanitize and validate request
func SanitizeRequest(ctx echo.Context, request interface{}) error {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}
	defer ctx.Request().Body.Close()

	sanBody, err := sanitize.SanitizeJSON(body)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	if err = json.Unmarshal(sanBody, request); err != nil {
		return err
	}

	return validate.StructCtx(ctx.Request().Context(), request)
}

var allowedImagesContentTypes = map[string]string{
	"image/bmp":                "bmp",
	"image/gif":                "gif",
	"image/png":                "png",
	"image/jpeg":               "jpeg",
	"image/jpg":                "jpg",
	"image/svg+xml":            "svg",
	"image/webp":               "webp",
	"image/tiff":               "tiff",
	"image/vnd.microsoft.icon": "ico",
}

func CheckImageFileContentType(fileContent []byte) (string, error) {
	contentType := http.DetectContentType(fileContent)

	extension, ok := allowedImagesContentTypes[contentType]
	if !ok {
		return "", errors.New("this content type is not allowed")
	}

	return extension, nil
}
