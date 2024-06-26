package vo

import "fmt"

type CreateUsuarioRequest struct {
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Documento string `json:"documento"`
	Endereco  string `json:"endereco"`
}
type UpdateUsuarioRequest struct {
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Documento string `json:"documento"`
	Endereco  string `json:"endereco"`
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *CreateUsuarioRequest) Validate() error {
	if r.Nome == "" && r.Telefone == "" && r.Email == "" && r.Documento == "" && r.Endereco == "" {
		return fmt.Errorf("request body is empty")
	}
	if r.Nome == "" {
		return ErrParamIsRequired("Nome", "string")
	}
	if r.Telefone == "" {
		return ErrParamIsRequired("Telefone", "string")
	}
	if r.Email == "" {
		return ErrParamIsRequired("Email", "string")
	}
	if r.Documento == "" {
		return ErrParamIsRequired("Documento", "string")
	}
	if r.Endereco == "" {
		return ErrParamIsRequired("Endereco", "string")
	}
	return nil
}
func (r *UpdateUsuarioRequest) Validate() error {
	if r.Nome == "" && r.Telefone == "" && r.Email == "" && r.Documento == "" && r.Endereco == "" {
		return fmt.Errorf("a least one field must be provied")
	}

	return nil
}
