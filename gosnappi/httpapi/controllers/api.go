package controllers

import (
	gosnappi "github.com/open-traffic-generator/snappi/gosnappi"
)

var globalApi *api

type api struct {
	api *gosnappi.GosnappiApi
}

type Api interface {
	getApi() gosnappi.GosnappiApi
}

func (obj *api) getApi() gosnappi.GosnappiApi {
	if obj.api == nil {
		return gosnappi.NewApi()
	}
	return *obj.api
}

func NewApi() *api {
	if globalApi == nil {
		globalApi = &api{}
		return globalApi
	}
	return globalApi
}
