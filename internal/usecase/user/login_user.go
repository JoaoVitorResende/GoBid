package user

import (
	"context"

	"github.com/JoaoVitorResende/GoBid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.Evaluator{
	var eval validator.Evaluator

	eval.Checkfield(validator.Matches(req.Email, validator.EmailRX), "email","must be a valid email")
	eval.Checkfield(validator.NotBlank(req.Password), "password", "this field cannot be blank")

	return eval
}
