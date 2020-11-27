package service

import (
	"context"
	"github.com/0ndreu/documents/internal/models"
)

type Service interface {
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

type storage interface {
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, ID int64) error
	CheckEmail(ctx context.Context, email string) (bool, error)

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

type service struct {
	storage storage
}

func (s *service) IsAlive(ctx context.Context) (errMaster, errSlave error) {
	panic("implement me")
}

// NewService creates a new instance of service
func NewService(storage storage) Service {
	return &service{
		storage: storage,
	}
}
