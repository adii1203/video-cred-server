package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/adii1203/video-cred/internals/storage"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("invalid input")
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

func (s *UserService) GetUserById(ctx context.Context, id pgtype.UUID) (storage.User, error) {
	if err := ctx.Err(); err != nil {
		return storage.User{}, fmt.Errorf("operation aborted: %w", err)
	}

	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.User{}, ErrUserNotFound
		}
		return storage.User{}, fmt.Errorf("cannot get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUserByClerkId(ctx context.Context, clerkId pgtype.Text) (storage.User, error) {
	if err := ctx.Err(); err != nil {
		return storage.User{}, fmt.Errorf("operation aborted: %w", err)
	}

	user, err := s.repo.GetUserByClerkId(ctx, clerkId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return storage.User{}, ErrUserNotFound
		}
		return storage.User{}, fmt.Errorf("cannot get user: %w", err)
	}

	return user, nil
}
