package consts

import "errors"

var (
	ErrInvalidIdCompany       = errors.New("id da empresa inválido")
	ErrInvalidRazaoSocial     = errors.New("razão social inválida")
	ErrInvalidNomeFantasia    = errors.New("nome fantasia inválido")
	ErrInvalidCnpj            = errors.New("cnpj inválido")
	ErrInvalidIe              = errors.New("ie inválido")
	ErrInvalidIm              = errors.New("im inválido")
	ErrInvalidEndereco        = errors.New("endereço inválido")
	ErrInvalidTelefone        = errors.New("telefone inválido")
	ErrInvalidEmail           = errors.New("email inválido")
	ErrInvalidSite            = errors.New("site inválido")
	ErrContatoPrincipal       = errors.New("contato principal inválido")
	ErrInvalidSegmentoAtuacao = errors.New("segmento de atuação inválido")
	ErrInvalidObservacao      = errors.New("observação inválida")
	ErrInvalidCreatedAt       = errors.New("data de criação inválida")
	ErrInvalidUpdatedAt       = errors.New("data de atualização inválida")
	ErrInvalidUserId          = errors.New("id do usuário inválido")

	ErrNotEmptyIdCompany        = errors.New("id da empresa não pode ser vazio")
	ErrNotEmptyRazaoSocial      = errors.New("razão social não pode ser vazio")
	ErrNotEmptyNomeFantasia     = errors.New("nome fantasia não pode ser vazio")
	ErrNotEmptyCnpj             = errors.New("cnpj não pode ser vazio")
	ErrNotEmptyIe               = errors.New("ie não pode ser vazio")
	ErrNotEmptyIm               = errors.New("im não pode ser vazio")
	ErrNotEmptyEndereco         = errors.New("endereço não pode ser vazio")
	ErrNotEmptyTelefone         = errors.New("telefone não pode ser vazio")
	ErrNotEmptyEmail            = errors.New("email não pode ser vazio")
	ErrNotEmptySite             = errors.New("site não pode ser vazio")
	ErrNotEmptyContatoPrincipal = errors.New("contato principal não pode ser vazio")
	ErrNotEmptySegmentoAtuacao  = errors.New("segmento de atuação não pode ser vazio")
	ErrNotEmptyObservacao       = errors.New("observação não pode ser vazio")
	ErrNotEmptyCreatedAt        = errors.New("data de criação não pode ser vazio")
	ErrNotEmptyUpdatedAt        = errors.New("data de atualização não pode ser vazio")
	ErrNotEmptyUserId           = errors.New("id do usuário não pode ser vazio")

	// errors not found
	ErrNotFoundCompany = errors.New("empresa não encontrado")

	// errors already exists
	ErrAlreadyExistsCompany = errors.New("empresa já cadastrado")

	// Maximum and minimum lengths
	MaximumLengthRazaoSocialCompany       = 100
	MaximumLengthNomeFantasiaCompany      = 100
	MaximumLengthCnpjCompany              = 14
	MaximumLengthIeCompany                = 20
	MaximumLengthAddressCompany           = 100
	MaximumLengthPhoneCompany             = 20
	MaximumLengthCellPhoneCompany         = 20
	MaximumLengthEmailCompany             = 100
	MaximumLengthSiteCompany              = 100
	MaximumLengthContactCompany           = 100
	MaximumLengthSegmentCompany           = 100
	MaximumLengthPaymentConditionsCompany = 100
	MaximumLengthObservationCompany       = 100
	MaximumLengthUserIdCompany            = 100

	MinimumLengthRazaoSocialCompany       = 3
	MinimumLengthNomeFantasiaCompany      = 3
	MinimumLengthCnpjCompany              = 11
	MinimumLengthIeCompany                = 9
	MinimumLengthAddressCompany           = 3
	MinimumLengthPhoneCompany             = 8
	MinimumLengthCellPhoneCompany         = 8
	MinimumLengthEmailCompany             = 3
	MinimumLengthSiteCompany              = 3
	MinimumLengthContactCompany           = 3
	MinimumLengthSegmentCompany           = 3
	MinimumLengthPaymentConditionsCompany = 3
	MinimumLengthObservationCompany       = 3
	MinimumLengthUserIdCompany            = 3

	ErrMaximumLengthrRazaoSocialCompany      = errors.New("razao social deve ter no maximo 100 caracteres")
	ErrMaximumLengthrNomeFantasiaCompany     = errors.New("nome fantasia deve ter no maximo 100 caracteres")
	ErrMaximumLengthCnpjCompany              = errors.New("o cnpj deve ter no máximo 14 caracteres")
	ErrMaximumLengthIeCompany                = errors.New("a i.e da empresa deve ter no máximo 20 caracteres")
	ErrMaximumLengthAddressCompany           = errors.New("o endereço da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthPhoneCompany             = errors.New("o telefone da empresa deve ter no máximo 20 caracteres")
	ErrMaximumLengthCellPhoneCompany         = errors.New("o celular da empresa deve ter no máximo 20 caracteres")
	ErrMaximumLengthEmailCompany             = errors.New("o email da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthSiteCompany              = errors.New("o site da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthContactCompany           = errors.New("o contato principal da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthSegmentCompany           = errors.New("o segmento de atuação da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthPaymentConditionsCompany = errors.New("as condições de pagamento da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthObservationCompany       = errors.New("a observação da empresa deve ter no máximo 100 caracteres")
	ErrMaximumLengthUserIdCompany            = errors.New("o id do usuário da empresa deve ter no máximo 100 caracteres")

	ErrMinimumLengthRazaoSocialCompany       = errors.New("razao social deve ter no minimo 3 caracteres")
	ErrMinimumLengthNomeFantasiaCompany      = errors.New("nome fantasia deve ter no minimo 3 caracteres")
	ErrMinimumLengthCnpjCompany              = errors.New("o cnpj da empresa deve ter no mínimo 11 caracteres")
	ErrMinimumLengthIeCompany                = errors.New("a i.e da empresa deve ter no mínimo 9 caracteres")
	ErrMinimumLengthAddressCompany           = errors.New("o endereço da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthPhoneCompany             = errors.New("o telefone da empresa deve ter no mínimo 8 caracteres")
	ErrMinimumLengthCellPhoneCompany         = errors.New("o celular da empresa deve ter no mínimo 8 caracteres")
	ErrMinimumLengthEmailCompany             = errors.New("o email da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthSiteCompany              = errors.New("o site da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthContactCompany           = errors.New("o contato principal da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthSegmentCompany           = errors.New("o segmento de atuação da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthPaymentConditionsCompany = errors.New("as condições de pagamento da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthObservationCompany       = errors.New("a observação da empresa deve ter no mínimo 3 caracteres")
	ErrMinimumLengthUserIdCompany            = errors.New("o id do usuário da empresa deve ter no mínimo 3 caracteres")
)
