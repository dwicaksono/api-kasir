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

func (u *categoryUsecase) GetAll(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.GetAll(ctx)
}

func (u *categoryUsecase) GetByID(ctx context.Context, id int) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.GetByID(ctx, id)
}

func (u *categoryUsecase) Create(ctx context.Context, c *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Create(ctx, c)
}

func (u *categoryUsecase) Update(ctx context.Context, id int, c *domain.Category) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Update(ctx, id, c)
}

func (u *categoryUsecase) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.categoryRepo.Delete(ctx, id)
}
