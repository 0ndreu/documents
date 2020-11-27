package server

import (
	"context"
	"github.com/0ndreu/documents/internal/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type service interface {
	IsAlive(ctx context.Context) (errMaster, errSlave error)

	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, ID int64) error
	CheckUser(ctx context.Context, user *models.User) (bool, error)

	CreateDocument(ctx context.Context, doc *models.Document) error
	UpdateDocument(ctx context.Context, doc *models.Document) error
	DeleteDocument(ctx context.Context, ID int64) error
	GetDocumentsList(ctx context.Context) ([]models.Document, error)
	GetDocumentByID(ctx context.Context, ID int64) (*models.Document, error)

	CreateClassifier(ctx context.Context, doc *models.Classifier) error
	UpdateClassifier(ctx context.Context, doc *models.Classifier) error
	DeleteClassifier(ctx context.Context, ID int32) error
	GetClassifiersList(ctx context.Context) ([]models.Classifier, error)
	GetClassifierByID(ctx context.Context, ID int32) (*models.Classifier, error)
}

// Server interface
type Server interface {
	CheckLiveness(w http.ResponseWriter, r *http.Request)
}

type middleware interface {
	IsValid(token string) (ok bool)
}

type server struct {
	ctx  context.Context
	svc  service
	log  zerolog.Logger
	auth middleware
}

// NewServer creates a new services router
func NewServer(ctx context.Context, svc service, log zerolog.Logger) http.Handler {
	srv := &server{
		ctx: ctx,
		svc: svc,
		log: log,
	}

	r := mux.NewRouter()
	sv1 := r.PathPrefix("/api/v1").Subrouter()
	//sv1.Use(srv.basicVerify)
	//sv1.Use(srv.recovery)
	//sv1.Use(srv.respondOptions)
	r.HandleFunc("/health", srv.CheckLiveness).Methods(http.MethodGet)

	sv1.HandleFunc("/user/new", srv.CreateUser).Methods(http.MethodPost)
	sv1.HandleFunc("/user/{id}/delete", srv.DeleteUser).Methods(http.MethodPost)
	sv1.HandleFunc("/user/{id}", srv.UpdateUser).Methods(http.MethodPost)
	sv1.HandleFunc("/user/auth", srv.AuthUser).Methods(http.MethodPost)

	sv1.HandleFunc("/document/new", srv.CreateDocument).Methods(http.MethodPost)
	sv1.HandleFunc("/document", srv.UpdateDocument).Methods(http.MethodPost)
	sv1.HandleFunc("/document/{id}/delete", srv.DeleteDocument).Methods(http.MethodPost)
	sv1.HandleFunc("/document/list", srv.GetDocumentsList).Methods(http.MethodGet)
	sv1.HandleFunc("/document/{id}", srv.GetDocumentByID).Methods(http.MethodGet)

	sv1.HandleFunc("/classifier/new", srv.CreateClassifier).Methods(http.MethodPost)
	sv1.HandleFunc("/classifier", srv.UpdateClassifier).Methods(http.MethodPost)
	sv1.HandleFunc("/classifier/{id}/delete", srv.DeleteClassifier).Methods(http.MethodPost)
	sv1.HandleFunc("/classifier/list", srv.GetClassifiersList).Methods(http.MethodGet)
	sv1.HandleFunc("/classifier/{id}", srv.GetClassifierByID).Methods(http.MethodGet)

	return r
}
