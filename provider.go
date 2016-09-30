package gosoc

import (
	"errors"

	"github.com/magnusfarje/gosoc/models"
)

// Provider - Provider interface
type Provider interface {
	Name() string
	ValidateToken(token string) (models.User, error)
}

// Providers - Collection of providers
type Providers map[string]Provider

var _providers = Providers{}

//AddProviders - Add providers to use
func AddProviders(providers ...Provider) {
	for _, provider := range providers {
		_providers[provider.Name()] = provider
	}
}

// GetProvider - Get provider by name
func GetProvider(name string) (Provider, error) {
	provider := _providers[name]
	if provider == nil {
		return nil, errors.New("Provider don't exist")
	}
	return provider, nil
}
