package service

import (
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

const emptyString = "novaluestring"

type CasdoorService struct {
	client *casdoorsdk.Client
}

func NewCasdoorService(client *casdoorsdk.Client) *CasdoorService {
	return &CasdoorService{client: client}
}

type EnforceInput struct {
	PermissionId         string
	Resource             string
	Action               string
	RelationObjectUserId string
	AccessToken          string
}

type EnforceResult struct {
	allowed bool
}

func (s *CasdoorService) Enforce(input *EnforceInput) (*EnforceResult, error) {
	if input.AccessToken == "" {
		return nil, fmt.Errorf("accessToken missing")
	}

	claims, err := s.client.ParseJwtToken(input.AccessToken)

	if err != nil {
		return nil, fmt.Errorf("error while parsing casdoor access token" + "; " + err.Error())
	}

	sub := claims.User.Owner + "/" + claims.User.Name

	casbinRequest := []interface{}{
		sub,
		input.Resource,
		input.Action,
		input.RelationObjectUserId,
		// claims.User.ExternalId,
	}

	result, err := s.client.Enforce(input.PermissionId, emptyString, emptyString, emptyString, emptyString, casbinRequest)

	if err != nil {
		return nil, err
	}

	return &EnforceResult{
		allowed: result,
	}, nil
}
