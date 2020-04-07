package controller

import "go.uber.org/dig"

// GatewayCtrl ...
type GatewayCtrl struct {
	dig.In
	CalculateCtrl *CalculateController
	ProblemCtrl   *ProblemController
	PeopleCtrl    *PeopleController
}
