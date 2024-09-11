package user

import (
	"context"
	"time"

	"vk-test-task/pkg/format"
	"vk-test-task/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Store interface {
		Create(context.Context, CreateEntity) (Entity, error)
		GetPassHashAndRoleByUsername(context.Context, string) (string, string, error)
		CheckExistence(context.Context, string) (bool, error)
		SaveRefreshToken(ctx context.Context, username string, refreshToken string) error
    	GetRefreshTokenByUsername(ctx context.Context, username string) (string, error)
    	UpdateRefreshToken(ctx context.Context, username string, refreshToken string) error
	}

	storeImpl struct {
		client *pgxpool.Pool
	}

	CreateEntity struct {
		Username string
		PassHash string
		Role     string
	}

	Entity struct {
		ID        int
		Username  string
		PassHash  string
		Role      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenEntity struct {
		Username    string
		RefreshToken string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)

func New(client *pgxpool.Pool) Store {
	return &storeImpl{
		client: client,
	}
}

func (s *storeImpl) Create(ctx context.Context, entity CreateEntity) (Entity, error) {
	var newUser Entity

	err := s.client.QueryRow(
		ctx,
		`
			INSERT INTO users (username, pass_hash, role, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, username, role, created_at, updated_at
		`,
		entity.Username,
		entity.PassHash,
		entity.Role,
		format.TimeNow(),
		format.TimeNow(),
	).Scan(&newUser.ID,
		&newUser.Username,
		&newUser.Role,
		&newUser.CreatedAt,
		&newUser.UpdatedAt)
	if err != nil {
		logger.Log.Error("create new user",
			"error", err.Error())
	}

	return newUser, err
}

func (s *storeImpl) GetPassHashAndRoleByUsername(ctx context.Context, username string) (string, string, error) {
	var passHash string
	var role string

	err := s.client.QueryRow(
		ctx,
		`
			SELECT pass_hash, role
			FROM users
			WHERE username = $1
		`,
		username,
	).Scan(&passHash, &role)
	if err != nil {
		logger.Log.Error("get pass hash and role by username",
			"error", err.Error())
	}

	return passHash, role, err
}

func (s *storeImpl) CheckExistence(ctx context.Context, username string) (bool, error) {
	var exists bool

	err := s.client.QueryRow(
		ctx,
		`
            SELECT EXISTS (
                SELECT 1
                FROM users
                WHERE username = $1
            )
        `,
		username,
	).Scan(&exists)
	if err != nil {
		logger.Log.Error("check user existence",
			"error", err.Error())
	}

	return exists, err
}

func (s *storeImpl) SaveRefreshToken(ctx context.Context, username string, refreshToken string) error {
    _, err := s.client.Exec(
        ctx,
        `
            INSERT INTO refresh_tokens (username, refresh_token, created_at, updated_at)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (username) DO UPDATE
            SET refresh_token = EXCLUDED.refresh_token,
                updated_at = EXCLUDED.updated_at
        `,
        username,
        refreshToken,
        format.TimeNow(),
        format.TimeNow(),
    )
    if err != nil {
        logger.Log.Error("save refresh token",
            "error", err.Error())
    }

    return err
}

func (s *storeImpl) GetRefreshTokenByUsername(ctx context.Context, username string) (string, error) {
    var refreshToken string

    err := s.client.QueryRow(
        ctx,
        `
            SELECT refresh_token
            FROM refresh_tokens
            WHERE username = $1
        `,
        username,
    ).Scan(&refreshToken)
    if err != nil {
        logger.Log.Error("get refresh token by username",
            "error", err.Error())
    }

    return refreshToken, err
}

func (s *storeImpl) UpdateRefreshToken(ctx context.Context, username string, refreshToken string) error {
    _, err := s.client.Exec(
        ctx,
        `
            UPDATE refresh_tokens
            SET refresh_token = $1,
                updated_at = $2
            WHERE username = $3
        `,
        refreshToken,
        format.TimeNow(),
        username,
    )
    if err != nil {
        logger.Log.Error("update refresh token",
            "error", err.Error())
    }

    return err
}