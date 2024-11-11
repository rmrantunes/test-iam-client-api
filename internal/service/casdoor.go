package service

import (
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

type CasdoorService struct {
	client *casdoorsdk.Client
}

func NewCasdoorService(client *casdoorsdk.Client) *CasdoorService {
	return &CasdoorService{client: client}
}

type EnforceInput struct {
	permissionId         string
	accessToken          string
	relationObjectUserId string
	resource             string
	action               string
}

type EnforceResult struct {
	allowed bool
}

func (s *CasdoorService) Enforce(input *EnforceInput) (*EnforceResult, error) {
	if input.accessToken == "" {
		return nil, fmt.Errorf("accessToken missing")
	}

	claims, err := s.client.ParseJwtToken(input.accessToken)

	if err != nil {
		return nil, fmt.Errorf("error while parsing casdoor access token")
	}

	sub := claims.User.Owner + "/" + claims.User.Name

	casbinRequest := []interface{}{
		sub,
		input.resource,
		input.action,
		input.relationObjectUserId,
		claims.User.ExternalId,
	}

	result, err := s.client.Enforce(input.permissionId, "", "", "", "", casbinRequest)

	return &EnforceResult{
		allowed: result,
	}, nil
}
