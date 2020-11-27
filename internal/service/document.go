package service

import (
	"context"
	"github.com/0ndreu/documents/internal/models"
)

func (s *service) CreateDocument(ctx context.Context, doc *models.Document) error {
	return s.storage.CreateDocument(ctx, doc)
}

func (s *service) UpdateDocument(ctx context.Context, doc *models.Document) error {
	return s.storage.UpdateDocument(ctx, doc)
}

func (s *service) DeleteDocument(ctx context.Context, ID int64) error {
	return s.storage.DeleteDocument(ctx, ID)
}

func (s *service) GetDocumentsList(ctx context.Context) ([]models.Document, error) {
	return s.storage.GetDocumentsList(ctx)
}

func (s *service) GetDocumentByID(ctx context.Context, ID int64) (*models.Document, error) {
	return s.storage.GetDocumentByID(ctx, ID)
}
