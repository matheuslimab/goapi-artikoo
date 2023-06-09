package consts

import "errors"

var (
	// errors invalids
	ErrInvalidIdClient                = errors.New("id do cliente inválido")
	ErrInvalidTypeClient              = errors.New("tipo de cliente inválido")
	ErrInvalidNameClient              = errors.New("nome do cliente inválido")
	ErrInvalidCpfCnpjClient           = errors.New("cpf ou cnpj do cliente inválido")
	ErrInvalidRgIeClient              = errors.New("rg ou ie do cliente inválido")
	ErrInvalidBirthDateClient         = errors.New("data de nascimento do cliente inválido")
	ErrInvalidAddressClient           = errors.New("endereço do cliente inválido")
	ErrInvalidPhoneClient             = errors.New("telefone do cliente inválido")
	ErrInvalidCellPhoneClient         = errors.New("celular do cliente inválido")
	ErrInvalidEmailClient             = errors.New("email do cliente inválido")
	ErrInvalidSiteClient              = errors.New("site do cliente inválido")
	ErrInvalidContactClient           = errors.New("contato principal do cliente inválido")
	ErrInvalidSegmentClient           = errors.New("segmento de atuação do cliente inválido")
	ErrInvalidCreditLimitClient       = errors.New("limite de crédito do cliente inválido")
	ErrInvalidPaymentConditionsClient = errors.New("condições de pagamento do cliente inválido")
	ErrInvalidStatusClient            = errors.New("status do cliente inválido")
	ErrInvalidObservationClient       = errors.New("observação do cliente inválido")
	ErrInvalidCreatedAtClient         = errors.New("data de criação do cliente inválido")
	ErrInvalidUpdateAtClient          = errors.New("data de atualização do cliente inválido")
	ErrInvalidCompanyIdClient         = errors.New("id da empresa do cliente inválido")
	ErrInvalidUserIdClient            = errors.New("id do usuário do cliente inválido")

	// errors not empty
	ErrNotEmptyIdClient                = errors.New("id do cliente não pode ser vazio")
	ErrNotEmptyTypeClient              = errors.New("tipo de cliente não pode ser vazio")
	ErrNotEmptyNameClient              = errors.New("nome do cliente não pode ser vazio")
	ErrNotEmptyCpfCnpjClient           = errors.New("cpf ou cnpj do cliente não pode ser vazio")
	ErrNotEmptyRgIeClient              = errors.New("rg ou ie do cliente não pode ser vazio")
	ErrNotEmptyBirthDateClient         = errors.New("data de nascimento do cliente não pode ser vazio")
	ErrNotEmptyAddressClient           = errors.New("endereço do cliente não pode ser vazio")
	ErrNotEmptyPhoneClient             = errors.New("telefone do cliente não pode ser vazio")
	ErrNotEmptyCellPhoneClient         = errors.New("celular do cliente não pode ser vazio")
	ErrNotEmptyEmailClient             = errors.New("email do cliente não pode ser vazio")
	ErrNotEmptySiteClient              = errors.New("site do cliente não pode ser vazio")
	ErrNotEmptyContactClient           = errors.New("contato principal do cliente não pode ser vazio")
	ErrNotEmptySegmentClient           = errors.New("segmento de atuação do cliente não pode ser vazio")
	ErrNotEmptyCreditLimitClient       = errors.New("limite de crédito do cliente não pode ser vazio")
	ErrNotEmptyPaymentConditionsClient = errors.New("condições de pagamento do cliente não pode ser vazio")
	ErrNotEmptyStatusClient            = errors.New("status do cliente não pode ser vazio")
	ErrNotEmptyObservationClient       = errors.New("observação do cliente não pode ser vazio")
	ErrNotEmptyCreatedAtClient         = errors.New("data de criação do cliente não pode ser vazio")
	ErrNotEmptyUpdateAtClient          = errors.New("data de atualização do cliente não pode ser vazio")
	ErrNotEmptyCompanyIdClient         = errors.New("id da empresa do cliente não pode ser vazio")
	ErrNotEmptyUserIdClient            = errors.New("id do usuário do cliente não pode ser vazio")

	// errors not found
	ErrNotFoundClient = errors.New("cliente não encontrado")

	// errors already exists
	ErrAlreadyExistsClient = errors.New("cliente já cadastrado")

	// Maximum and minimum lengths
	MaximumLengthNameClient              = 100
	MaximumLengthCpfCnpjClient           = 14
	MaximumLengthRgIeClient              = 20
	MaximumLengthAddressClient           = 100
	MaximumLengthPhoneClient             = 20
	MaximumLengthCellPhoneClient         = 20
	MaximumLengthEmailClient             = 100
	MaximumLengthSiteClient              = 100
	MaximumLengthContactClient           = 100
	MaximumLengthSegmentClient           = 100
	MaximumLengthPaymentConditionsClient = 100
	MaximumLengthObservationClient       = 100
	MaximumLengthCompanyIdClient         = 100
	MaximumLengthUserIdClient            = 100

	ErrMaximumLengthNameClient              = errors.New("o nome da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthCpfCnpjClient           = errors.New("o cpf/cnpj da empresa/cliente deve ter no máximo 14 caracteres")
	ErrMaximumLengthRgIeClient              = errors.New("o rg/ie da empresa/cliente deve ter no máximo 20 caracteres")
	ErrMaximumLengthAddressClient           = errors.New("o endereço da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthPhoneClient             = errors.New("o telefone da empresa/cliente deve ter no máximo 20 caracteres")
	ErrMaximumLengthCellPhoneClient         = errors.New("o celular da empresa/cliente deve ter no máximo 20 caracteres")
	ErrMaximumLengthEmailClient             = errors.New("o email da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthSiteClient              = errors.New("o site da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthContactClient           = errors.New("o contato principal da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthSegmentClient           = errors.New("o segmento de atuação da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthPaymentConditionsClient = errors.New("as condições de pagamento da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthObservationClient       = errors.New("a observação da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthCompanyIdClient         = errors.New("o id da empresa da empresa/cliente deve ter no máximo 100 caracteres")
	ErrMaximumLengthUserIdClient            = errors.New("o id do usuário da empresa/cliente deve ter no máximo 100 caracteres")

	MinimumLengthNameClient              = 3
	MinimumLengthCpfCnpjClient           = 11
	MinimumLengthRgIeClient              = 9
	MinimumLengthAddressClient           = 3
	MinimumLengthPhoneClient             = 8
	MinimumLengthCellPhoneClient         = 8
	MinimumLengthEmailClient             = 3
	MinimumLengthSiteClient              = 3
	MinimumLengthContactClient           = 3
	MinimumLengthSegmentClient           = 3
	MinimumLengthPaymentConditionsClient = 3
	MinimumLengthObservationClient       = 3
	MinimumLengthCompanyIdClient         = 3
	MinimumLengthUserIdClient            = 3

	ErrMinimumLengthNameClient              = errors.New("o nome da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthCpfCnpjClient           = errors.New("o cpf/cnpj da empresa/cliente deve ter no mínimo 11 caracteres")
	ErrMinimumLengthRgIeClient              = errors.New("o rg/ie da empresa/cliente deve ter no mínimo 9 caracteres")
	ErrMinimumLengthAddressClient           = errors.New("o endereço da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthPhoneClient             = errors.New("o telefone da empresa/cliente deve ter no mínimo 8 caracteres")
	ErrMinimumLengthCellPhoneClient         = errors.New("o celular da empresa/cliente deve ter no mínimo 8 caracteres")
	ErrMinimumLengthEmailClient             = errors.New("o email da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthSiteClient              = errors.New("o site da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthContactClient           = errors.New("o contato principal da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthSegmentClient           = errors.New("o segmento de atuação da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthPaymentConditionsClient = errors.New("as condições de pagamento da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthObservationClient       = errors.New("a observação da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthCompanyIdClient         = errors.New("o id da empresa da empresa/cliente deve ter no mínimo 3 caracteres")
	ErrMinimumLengthUserIdClient            = errors.New("o id do usuário da empresa/cliente deve ter no mínimo 3 caracteres")

	// API ERRORS

	ErrAuthLoginAPI  = errors.New("operação não autorizada, talvez seus dados de autenticação não estejam certos.")
	ErrAuthApiKeyAPI = errors.New("operação não autorizada, talvez sua  api key esteja incorreta.")
)
