package server

import "time"

/*
 * Package for server authentication and authorization
 *
 * Implements JWT-based authentication and authorization
 */

// Guest will be the default role for simplifying the development
// of the RBAC system
const (
	RoleGuest = 0
	RoleUser  = 1
	RoleAdmin = 1000
)

// We don't bother with email
type Account struct {
	ID       string    `json:"id,omitempty,nonempty" toml:"id"`
	UUID     string    `json:"uuid,omitempty" toml:"uuid"`
	Password string    `json:"password,omitempty" toml:"password"`
	Creation time.Time `json:"creation,omitempty" toml:"creation"`
	Role     int       `json:"role,omitempty" xml:"role,omitempty" toml:"role"`
}
