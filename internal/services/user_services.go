package services

import (
	"context"
	"errors"

	"github.com/JoaoVitorResende/GoBid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// pool serve para manter o banco de dados aberto igual ao ado.net
type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

var (
	ErrDuplicatedEmailOrUserName = errors.New("User name or email already exists")
	ErrInvalidCredentials = errors.New("Login invalid")
)

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: hash,
		Bio:          bio,
	}

	id, err := us.queries.CreateUser(ctx, args)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUserName
		}
		return uuid.UUID{}, err
	}
	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil{
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword){
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}
	return user.ID, nil
}
