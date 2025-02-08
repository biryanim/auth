package access

import (
	"context"
	"errors"
	"fmt"
	"github.com/biryanim/auth/internal/utils"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authPrefix = "Bearer "
)

var accessibleRoles map[string]int32

func (s *serv) accessibleRoles(ctx context.Context) (map[string]int32, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]int32)

		accessInfo, err := s.accessRepository.GetList(ctx)
		if err != nil {
			fmt.Println("\n\nasdfafasfasfasf", err, accessInfo)
			return nil, err
		}

		for _, info := range accessInfo {
			accessibleRoles[info.EndpointAddress] = info.Role
		}
	}

	return accessibleRoles, nil
}

func (s *serv) Check(ctx context.Context, endpointAddress string) (bool, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false, errors.New("metadata not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return false, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return false, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, s.authConfig.AccessTokenSecret())
	if err != nil {
		return false, errors.New("access token is invalid")
	}

	accessMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return false, errors.New("failed to get accessible roles")
	}

	role, ok := accessMap[endpointAddress]
	if !ok {
		return true, nil
	}

	if role == claims.Role {
		return true, nil
	}

	return false, errors.New("access denied")
}
