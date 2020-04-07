package repository

import "go.uber.org/dig"

// GatewayRepo ...
type GatewayRepo struct {
	dig.In
	PeopleRepo *PeopleRepository
}
