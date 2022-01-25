package services

import (
	"go-viacep-wrapper/internal/application/gateway/viacep"
	domain "go-viacep-wrapper/internal/domain/models"
)

type ProfileService interface {
	FindAddress(zipcode string) (*domain.Address, error)
}

type addressService struct {
	api viacep.Gateway
}

func NewProfileService(api viacep.Gateway) ProfileService {
	return &addressService{api: api}
}

// FindAddress finds a address by passing an zipcode
func (service *addressService) FindAddress(zipcode string) (*domain.Address, error) {
	dto, err := service.api.GetAddress(zipcode)
	if err != nil {
		return nil, err
	}
	return hidrateAddress(dto), nil
}

func hidrateAddress(dto *viacep.ViaCepResponse) *domain.Address {
	output := &domain.Address{}
	output.AddressLine1 = dto.Logradouro
	output.AddressLine2 = dto.Bairro
	output.AddressLine3 = dto.Complemento
	output.City = dto.Localidade
	output.State = dto.Uf
	output.Country = "Brazil"
	output.ZipCode = dto.Cep
	return output
}
