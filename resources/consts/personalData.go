package consts

import "errors"

type PersonalData struct {
	IdPersonalData  string `json:"id_dataUser"`
	Name            string `json:"name_dataUser"`
	Email           string `json:"email_dataUser"`
	Cpf             string `json:"cpf_dataUser"`
	Rg              string `json:"rg_dataUser"`
	DataNascimento  string `json:"dataNascimento_dataUser"`
	Cargo           string `json:"cargo_dataUser"`
	TimeDeparture   string `json:"timeDeparture_DataUser"`
	AboutMe         string `json:"aboutMe_dataUser"`
	Ausencia        string `json:"ausencia_dataUser"`
	Genero          string `json:"genero_dataUser"`
	EnderecoPessoal string `json:"endereco_dataUser"`
	IdUser          string `json:"idUser_dataUser"`
}

var (

	// PersonalData campos inválidos
	ErrInvalidIdPersonalData              = errors.New("id inválido")
	ErrInvalidNamePersonalData            = errors.New("nome inválido")
	ErrInvalidEmailPersonalData           = errors.New("email inválido")
	ErrInvalidCpfPersonalData             = errors.New("cpf inválido")
	ErrInvalidRgPersonalData              = errors.New("rg inválido")
	ErrInvalidDataNascimentoPersonalData  = errors.New("data de nascimento inválida")
	ErrInvalidCargoPersonalData           = errors.New("cargo inválido")
	ErrInvalidTimeDeparturePersonalData   = errors.New("equipe/Departamento inválido")
	ErrInvalidAboutMePersonalData         = errors.New("sobre mim inválido")
	ErrInvalidAusenciaPersonalData        = errors.New("ausência inválida")
	ErrInvalidGeneroPersonalData          = errors.New("gênero inválido")
	ErrInvalidEnderecoPessoalPersonalData = errors.New("endereço inválido")
	ErrInvalidIdUserPersonalData          = errors.New("id do usuário inválido")

	// errors not empty
	ErrNotEmptyIdPersonalData              = errors.New("id não pode ser vazio")
	ErrNotEmptyNamePersonalData            = errors.New("nome não pode ser vazio")
	ErrNotEmptyEmailPersonalData           = errors.New("email não pode ser vazio")
	ErrNotEmptyCpfPersonalData             = errors.New("cpf não pode ser vazio")
	ErrNotEmptyRgPersonalData              = errors.New("rg não pode ser vazio")
	ErrNotEmptyDataNascimentoPersonalData  = errors.New("data de nascimento não pode ser vazio")
	ErrNotEmptyCargoPersonalData           = errors.New("cargo não pode ser vazio")
	ErrNotEmptyTimeDeparturePersonalData   = errors.New("equipe/Departamento não pode ser vazio")
	ErrNotEmptyAboutMePersonalData         = errors.New("sobre mim não pode ser vazio")
	ErrNotEmptyAusenciaPersonalData        = errors.New("ausência não pode ser vazio")
	ErrNotEmptyGeneroPersonalData          = errors.New("gênero não pode ser vazio")
	ErrNotEmptyEnderecoPessoalPersonalData = errors.New("endereço não pode ser vazio")
	ErrNotEmptyIdUserPersonalData          = errors.New("id do usuário não pode ser vazio")

	// errors not found
	ErrNotFoundPersonalData = errors.New("dados do usuário não encontrado")

	// Maximum lengths
	ErrMaxLengthNamePersonalData            = errors.New("nome deve ter no máximo 50 caracteres")
	ErrMaxLengthEmailPersonalData           = errors.New("email deve ter no máximo 100 caracteres")
	ErrMaxLengthCpfPersonalData             = errors.New("cpf deve ter no máximo 11 caracteres")
	ErrMaxLengthRgPersonalData              = errors.New("rg deve ter no máximo 10 caracteres")
	ErrMaxLengthDataNascimentoPersonalData  = errors.New("data de nascimento deve ter no máximo 10 caracteres")
	ErrMaxLengthCargoPersonalData           = errors.New("cargo deve ter no máximo 50 caracteres")
	ErrMaxLengthTimeDeparturePersonalData   = errors.New("equipe/Departamento deve ter no máximo 100 caracteres")
	ErrMaxLengthAboutMePersonalData         = errors.New("sobre mim deve ter no máximo 1000 caracteres")
	ErrMaxLengthAusenciaPersonalData        = errors.New("ausência deve ter no máximo 100 caracteres")
	ErrMaxLengthGeneroPersonalData          = errors.New("gênero deve ter no máximo 24 caracteres")
	ErrMaxLengthEnderecoPessoalPersonalData = errors.New("endereço deve ter no máximo 255 caracteres")
	ErrMaxLengthIdUserPersonalData          = errors.New("id do usuário deve ter no máximo 255 caracteres")

	MaximumLengthIdPersonalData              = 255
	MaximumLengthNamePersonalData            = 50
	MaximumLengthEmailPersonalData           = 100
	MaximumLengthCpfPersonalData             = 20
	MaximumLengthRgPersonalData              = 20
	MaximumLengthDataNascimentoPersonalData  = 10
	MaximumLengthCargoPersonalData           = 50
	MaximumLengthTimeDeparturePersonalData   = 100
	MaximumLengthAboutMePersonalData         = 1000
	MaximumLengthAusenciaPersonalData        = 100
	MaximumLengthGeneroPersonalData          = 24
	MaximumLengthEnderecoPessoalPersonalData = 255
	MaximumLengthIdUserPersonalData          = 255

	// Minimum lengths
	ErrMinLengthNamePersonalData            = errors.New("nome deve ter no mínimo 3 caracteres")
	ErrMinLengthEmailPersonalData           = errors.New("email deve ter no mínimo 5 caracteres")
	ErrMinLengthCpfPersonalData             = errors.New("cpf deve ter no mínimo 8 caracteres")
	ErrMinLengthRgPersonalData              = errors.New("rg deve ter no mínimo 5 caracteres")
	ErrMinLengthDataNascimentoPersonalData  = errors.New("data de nascimento deve ter no mínimo 3 caracteres")
	ErrMinLengthCargoPersonalData           = errors.New("cargo deve ter no mínimo 3 caracteres")
	ErrMinLengthTimeDeparturePersonalData   = errors.New("equipe/Departamento deve ter no mínimo 3 caracteres")
	ErrMinLengthAboutMePersonalData         = errors.New("sobre mim deve ter no mínimo 3 caracteres")
	ErrMinLengthGeneroPersonalData          = errors.New("gênero deve ter no mínimo 2 caracteres")
	ErrMinLengthEnderecoPessoalPersonalData = errors.New("endereço deve ter no mínimo 3 caracteres")
	ErrMinLengthIdUserPersonalData          = errors.New("id do usuário deve ter no mínimo 1 caracteres")

	MinimumLengthNamePersonalData            = 3
	MinimumLengthEmailPersonalData           = 5
	MinimumLengthCpfPersonalData             = 14
	MinimumLengthRgPersonalData              = 10
	MinimumLengthDataNascimentoPersonalData  = 3
	MinimumLengthCargoPersonalData           = 3
	MinimumLengthTimeDeparturePersonalData   = 3
	MinimumLengthAboutMePersonalData         = 3
	MinimumLengthGeneroPersonalData          = 2
	MinimumLengthEnderecoPessoalPersonalData = 3
	MinimumLengthIdUserPersonalData          = 1
)
