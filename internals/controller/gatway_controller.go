package controller

import "go.uber.org/dig"

type GatewayCtrl struct {
	dig.In
	PingCtrl *PingController
	// AuthorCtrl  *AuthorController
	ProductCtrl *ProductController
}
