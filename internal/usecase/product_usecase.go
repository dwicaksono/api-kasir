package usecase

import (
	"context"
	"kasir-api/internal/domain"
	"time"
)

type productUsecase struct {
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductUsecase(p domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo:    p,
		contextTimeout: timeout,
	}
}

func (u *productUsecase) GetAll(ctx context.Context, name string) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.productRepo.GetAll(ctx, name)
}

func (u *productUsecase) GetByID(ctx context.Context, id int) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.productRepo.GetByID(ctx, id)
}

func (u *productUsecase) Create(ctx context.Context, p *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.productRepo.Create(ctx, p)
}

func (u *productUsecase) Update(ctx context.Context, id int, p *domain.Product) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.productRepo.Update(ctx, id, p)
}

func (u *productUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.productRepo.Delete(ctx, id)
}
