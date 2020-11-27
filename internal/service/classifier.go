package service

import (
	"context"

	"github.com/0ndreu/documents/internal/models"
)

func (s *service) CreateClassifier(ctx context.Context, doc *models.Classifier) error {
	return s.storage.CreateClassifier(ctx, doc)
}

func (s *service) UpdateClassifier(ctx context.Context, doc *models.Classifier) error {
	return s.storage.UpdateClassifier(ctx, doc)
}

func (s *service) DeleteClassifier(ctx context.Context, ID int32) error {
	return s.storage.DeleteClassifier(ctx, ID)
}

func (s *service) GetClassifiersList(ctx context.Context) ([]models.Classifier, error) {
	return s.storage.GetClassifiersList(ctx)
}

func (s *service) GetClassifierByID(ctx context.Context, ID int32) (*models.Classifier, error) {
	return s.storage.GetClassifierByID(ctx, ID)
}
