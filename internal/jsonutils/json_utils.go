package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JoaoVitorResende/GoBid/internal/validator"
)
//transforma uma struct em JSON e envia como resposta, com o status HTTP.
func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-type", "Application")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encodejson %w", err)
	}

	return nil
}
// faz o mesmo que decodeJson e depois chama o método Valid() da struct para checar
//  as regras (nome vazio, email inválido etc.). Se algo estiver errado, devolve a lista de problemas.
func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var data T

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, nil, fmt.Errorf("decode json %w", err)
	}

	if problems := data.Valid(r.Context()); len(problems) > 0 {
		return data, problems, fmt.Errorf("invalid %T: %d problems", data, len(problems))
	}

	return data, nil, nil
}
//pega o JSON que chegou na requisição e transforma numa struct.
func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json failed: w%", err)
	}

	return data, nil
}
