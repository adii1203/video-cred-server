package service

import (
	"context"
	"fmt"

	"github.com/adii1203/video-cred/internals/storage"
)

type UserService struct {
	repo *storage.Queries
}

func NewUserService(repo *storage.Queries) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUserWithClerk(ctx context.Context, pramps storage.CreateUserParams) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("context error befour user created: %w", err)
	}

	_, err := s.repo.CreateUser(ctx, storage.CreateUserParams{
		Name:    pramps.Name,
		Email:   pramps.Email,
		Clerkid: pramps.Clerkid,
	})

	if err != nil {
		return fmt.Errorf("repository error creating user: %w", err)
	}

	return nil
}
