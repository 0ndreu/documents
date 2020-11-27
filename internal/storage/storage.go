package storage

import (
	"context"
	"github.com/0ndreu/documents/internal/models"
	"time"
)

type Storage interface {
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

type storage struct {
	clientUser     Client
	clientDocument Client
	timeout        time.Duration
}

func (s *storage) GetClassifiersList(ctx context.Context) ([]models.Classifier, error) {
	panic("implement me")
}

func (s *storage) CheckEmail(ctx context.Context, email string) (bool, error) {
	panic("implement me")
}

func (s *storage) GetDocumentsList(ctx context.Context) ([]models.Document, error) {
	panic("implement me")
}

func (s *storage) GetClassifiersByID(ctx context.Context, ID int64) (*models.Classifier, error) {
	panic("implement me")
}

func (s *storage) CreateUser(ctx context.Context, user *models.User) error {
	conn, _ := s.clientUser.GetClient()
	rows, err := conn.Query()
}

func (s *storage) UpdateUser(ctx context.Context, user *models.User) error {
	conn, _ := s.clientUser.GetClient()
	rows, err := conn.Query()
}

func (s *storage) DeleteUser(ctx context.Context, userID int64) error {
	conn, _ := s.clientUser.GetClient()
	rows, err := conn.Query()
}

func (s *storage) CheckUser(ctx context.Context, user *models.User) (bool, error) {
	conn, _ := s.clientUser.GetClient()
	rows, err := conn.Query()
}

func (s *storage) CreateDocument(ctx context.Context, document *models.Document) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) UpdateDocument(ctx context.Context, document *models.Document) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) DeleteDocument(ctx context.Context, documentID int64) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) GetDocumentList(ctx context.Context, document *models.Document) ([]models.Document, error) {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) GetDocumentByID(ctx context.Context, documentID int64) (*models.Document, error) {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) CreateClassifier(ctx context.Context, classifier *models.Classifier) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) UpdateClassifier(ctx context.Context, classifier *models.Classifier) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) DeleteClassifier(ctx context.Context, classifierID int32) error {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

func (s *storage) GetClassifierByID(ctx context.Context, classifierID int32) (*models.Classifier, error) {
	conn, _ := s.clientDocument.GetClient()
	rows, err := conn.Query()
}

// NewStorage creates a new instance of storage.
func NewStorage(clientUser, clientDocument Client, timeout time.Duration) Storage {
	return &storage{
		clientUser:     clientUser,
		clientDocument: clientDocument,
		timeout:        timeout,
	}
}
