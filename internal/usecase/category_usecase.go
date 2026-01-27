package usecase

import (
	"context"
	"kasir-api/internal/domain"
	"time"
)

type categoryUsecase struct {
	categoryRepo   domain.CategoryRepository
	contextTimeout time.Duration
}

func NewCategoryUsecase(c domain.CategoryRepository, timeout time.Duration) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryRepo:   c,
		contextTimeout: timeout,
	}
}

func (u *categoryUsecase) Fetch(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Fetch(ctx)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id int) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.GetByID(ctx, id)
}

func (u *categoryUsecase) Store(ctx context.Context, c *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Store(ctx, c)
}

func (u *categoryUsecase) Update(ctx context.Context, c *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Update(ctx, c)
}

func (u *categoryUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Delete(ctx, id)
}
