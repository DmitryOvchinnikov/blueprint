// Package crud binds the crud domain set of routes into the specified app.
package crud

import (
	"github.com/dmitryovchinnikov/blueprint/app/services/mold-api/v1/handlers/checkgrp"
	"github.com/dmitryovchinnikov/blueprint/app/services/mold-api/v1/handlers/homegrp"
	"github.com/dmitryovchinnikov/blueprint/app/services/mold-api/v1/handlers/productgrp"
	"github.com/dmitryovchinnikov/blueprint/app/services/mold-api/v1/handlers/trangrp"
	"github.com/dmitryovchinnikov/blueprint/app/services/mold-api/v1/handlers/usergrp"
	v1 "github.com/dmitryovchinnikov/blueprint/business/web/v1"
	"github.com/dmitryovchinnikov/blueprint/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg v1.APIMuxConfig) {
	checkgrp.Routes(app, checkgrp.Config{
		UsingWeaver: cfg.UsingWeaver,
		Build:       cfg.Build,
		Log:         cfg.Log,
		DB:          cfg.DB,
	})

	homegrp.Routes(app, homegrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})

	productgrp.Routes(app, productgrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})

	trangrp.Routes(app, trangrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})

	usergrp.Routes(app, usergrp.Config{
		Log:  cfg.Log,
		Auth: cfg.Auth,
		DB:   cfg.DB,
	})
}
