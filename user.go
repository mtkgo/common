package common

import "github.com/golang-jwt/jwt/v5"

type Permission struct {
	Name     string `json:"name,omitempty"`
	ExpireAt int64  `json:"exp,omitempty"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserID      string       `json:"userid"`
	Name        string       `json:"name,omitempty"` // only used for display, so use user.nickname
	AvatarURL   string       `json:"avatar_url,omitempty"`
	Roles       []string     `json:"roles,omitempty"` // like default springboot settings, has prefix "ROLE_"
	Permissions []Permission `json:"permissions,omitempty"`
}
