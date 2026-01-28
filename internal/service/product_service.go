package service

import (
	"context"
	"kasir-api/internal/domain"
)

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) GetByID(ctx context.Context, id int) (domain.Product, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, product *domain.Product) error {
	return s.repo.Create(ctx, product)
}

func (s *ProductService) Update(ctx context.Context, id int, product *domain.Product) error {
	return s.repo.Update(ctx, id, product)
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
