package service

import (
	"context"
	"kasir-api/internal/domain"
)

type CategoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *CategoryService) GetByID(ctx context.Context, id int) (domain.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) Create(ctx context.Context, category *domain.Category) error {
	return s.repo.Create(ctx, category)
}

func (s *CategoryService) Update(ctx context.Context, id int, category *domain.Category) error {
	return s.repo.Update(ctx, id, category)
}

func (s *CategoryService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
