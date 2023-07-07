package models

import (
	"strings"

	"github.com/matheuslimab/artikoo/api/resources/consts"
)

type PersonalData struct {
	IdPersonalData  string `json:"idPersonalData_dataUser"`
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
	Created_At      string `json:"created_at_dataUser"`
	Updated_At      string `json:"updated_at_dataUser"`
}

func (p *PersonalData) CheckData() error {
	err := p.validate()
	if err != nil {
		return err
	}

	err = p.formatar()
	if err != nil {
		return err
	}

	return nil
}

func (p *PersonalData) formatar() error {
	p.Name = strings.TrimSpace(p.Name)
	p.Email = strings.TrimSpace(p.Email)
	p.Cpf = strings.TrimSpace(p.Cpf)
	p.Rg = strings.TrimSpace(p.Rg)
	p.DataNascimento = strings.TrimSpace(p.DataNascimento)
	p.Cargo = strings.TrimSpace(p.Cargo)
	p.TimeDeparture = strings.TrimSpace(p.TimeDeparture)
	p.AboutMe = strings.TrimSpace(p.AboutMe)
	p.Ausencia = strings.TrimSpace(p.Ausencia)
	p.Genero = strings.TrimSpace(p.Genero)
	p.EnderecoPessoal = strings.TrimSpace(p.EnderecoPessoal)

	return nil
}

func (p *PersonalData) validate() error {

	// maximum

	if len(p.Name) > consts.MaximumLengthNamePersonalData {
		return consts.ErrMaxLengthNamePersonalData
	}

	if len(p.Email) > consts.MaximumLengthEmailPersonalData {
		return consts.ErrMaxLengthEmailPersonalData
	}

	if len(p.Cpf) > consts.MaximumLengthCpfPersonalData {
		return consts.ErrMaxLengthCpfPersonalData
	}

	if len(p.Rg) > consts.MaximumLengthRgPersonalData {
		return consts.ErrMaxLengthRgPersonalData
	}

	if len(p.DataNascimento) > consts.MaximumLengthDataNascimentoPersonalData {
		return consts.ErrMaxLengthDataNascimentoPersonalData
	}

	if len(p.Cargo) > consts.MaximumLengthCargoPersonalData {
		return consts.ErrMaxLengthCargoPersonalData
	}

	if len(p.TimeDeparture) > consts.MaximumLengthTimeDeparturePersonalData {
		return consts.ErrMaxLengthTimeDeparturePersonalData
	}

	if len(p.AboutMe) > consts.MaximumLengthAboutMePersonalData {
		return consts.ErrMaxLengthAboutMePersonalData
	}

	if len(p.Ausencia) > consts.MaximumLengthAusenciaPersonalData {
		return consts.ErrMaxLengthAusenciaPersonalData
	}

	if len(p.Genero) > consts.MaximumLengthGeneroPersonalData {
		return consts.ErrMaxLengthGeneroPersonalData
	}

	if len(p.EnderecoPessoal) > consts.MaximumLengthEnderecoPessoalPersonalData {
		return consts.ErrMaxLengthEnderecoPessoalPersonalData
	}

	// minimum

	if len(p.Name) < consts.MinimumLengthNamePersonalData {
		return consts.ErrMinLengthNamePersonalData
	}

	if len(p.Email) < consts.MinimumLengthEmailPersonalData {
		return consts.ErrMinLengthEmailPersonalData
	}

	if len(p.Cpf) < consts.MinimumLengthCpfPersonalData {
		return consts.ErrMinLengthCpfPersonalData
	}

	if len(p.Rg) < consts.MinimumLengthRgPersonalData {
		return consts.ErrMinLengthRgPersonalData
	}

	if len(p.DataNascimento) < consts.MinimumLengthDataNascimentoPersonalData {
		return consts.ErrMinLengthDataNascimentoPersonalData
	}

	if len(p.Cargo) < consts.MinimumLengthCargoPersonalData {
		return consts.ErrMinLengthCargoPersonalData
	}

	if len(p.TimeDeparture) < consts.MinimumLengthTimeDeparturePersonalData {
		return consts.ErrMinLengthTimeDeparturePersonalData
	}

	if len(p.AboutMe) < consts.MinimumLengthAboutMePersonalData {
		return consts.ErrMinLengthAboutMePersonalData
	}

	if len(p.EnderecoPessoal) < consts.MinimumLengthEnderecoPessoalPersonalData {
		return consts.ErrMinLengthEnderecoPessoalPersonalData
	}

	return nil
}
