package user

import (
	"context"

	"github.com/JoaoVitorResende/GoBid/internal/validator"
)

type CreateUserReq struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
	Bio          string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator{
	var eval validator.Evaluator
	eval.Checkfield(validator.NotBlank(req.UserName), "user_name", "this field cannot be empty")
	return eval
}
