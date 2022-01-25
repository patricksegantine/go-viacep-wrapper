package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type Gateway interface {
	GetAddress(zipcode string) (*ViaCepResponse, error)
}

func NewGateway(client *http.Client, host string) Gateway {
	return &apiGateway{
		client: client,
		host:   host,
	}
}

type apiGateway struct {
	client *http.Client
	host   string
}

// GetAddress query the `viacep` HTTP API for given zipcode
func (api *apiGateway) GetAddress(zipcode string) (*ViaCepResponse, error) {
	response, err := api.client.Get(fmt.Sprintf("%s/ws/%s/json", api.host, zipcode))
	if err != nil {
		return nil, err
	}
	bytes, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("there was an error getting the zip code")
	}

	// var output ViaCepResponse
	// err = json.NewDecoder(response.Body).Decode(&output)

	output := &ViaCepResponse{}
	err = json.Unmarshal(bytes, output)
	if err != nil {
		return nil, err
	}

	if output.Cep == "" {
		return nil, errors.New("zip code not found")
	}
	return output, nil
}
