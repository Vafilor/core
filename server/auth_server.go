package server

import (
	"context"
	"github.com/onepanelio/core/api"
	v1 "github.com/onepanelio/core/pkg"
	"github.com/onepanelio/core/server/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer contains logic for checking Authorization of resources in the system
type AuthServer struct {
}

// NewAuthServer creates a new AuthServer
func NewAuthServer() *AuthServer {
	return &AuthServer{}
}

// IsAuthorized checks if the provided action is authorized.
// No token == unauthorized. This is indicated by a nil ctx.
// Invalid token == unauthorized.
// Otherwise, we check with k8s using all of the provided data in the request.
func (a *AuthServer) IsAuthorized(ctx context.Context, request *api.IsAuthorizedRequest) (res *api.IsAuthorizedResponse, err error) {
	res = &api.IsAuthorizedResponse{}
	if ctx == nil {
		res.Authorized = false
		return res, status.Error(codes.Unauthenticated, "Unauthenticated.")
	}
	//User auth check
	client := getClient(ctx)

	err = a.isValidToken(err, client)
	if err != nil {
		return nil, err
	}

	//Check the request
	allowed, err := auth.IsAuthorized(client, request.IsAuthorized.Namespace, request.IsAuthorized.Verb, request.IsAuthorized.Group, request.IsAuthorized.Resource, request.IsAuthorized.ResourceName)
	if err != nil {
		res.Authorized = false
		return res, err
	}

	res.Authorized = allowed
	return res, nil
}

func (a *AuthServer) IsValidToken(ctx context.Context, req *api.IsValidTokenRequest) (res *api.IsValidTokenResponse, err error) {
	if ctx == nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated.")
	}

	client := getClient(ctx)

	err = a.isValidToken(err, client)
	if err != nil {
		return nil, err
	}

	config, err := client.GetSystemConfig()
	if err != nil {
		return
	}
	res = &api.IsValidTokenResponse{}
	res.Domain = config["ONEPANEL_DOMAIN"]

	return res, nil
}

func (a *AuthServer) isValidToken(err error, client *v1.Client) error {
	namespaces, err := client.ListOnepanelEnabledNamespaces()
	if err != nil {
		if err.Error() == "Unauthorized" {
			return status.Error(codes.Unauthenticated, "Unauthenticated.")
		}
		return err
	}
	if len(namespaces) == 0 {
		return errors.New("No namespaces for onepanel setup.")
	}
	namespace := namespaces[0]

	allowed, err := auth.IsAuthorized(client, "", "get", "", "namespaces", namespace.Name)
	if err != nil {
		return err
	}

	if !allowed {
		return status.Error(codes.Unauthenticated, "Unauthenticated.")
	}
	return nil
}
