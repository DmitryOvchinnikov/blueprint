package usersummarygrp

import (
	"net/http"

	"github.com/dmitryovchinnikov/blueprint/business/core/usersummary"
	"github.com/dmitryovchinnikov/blueprint/business/core/usersummary/stores/usersummarydb"
	"github.com/dmitryovchinnikov/blueprint/business/web/v1/auth"
	"github.com/dmitryovchinnikov/blueprint/business/web/v1/mid"
	"github.com/dmitryovchinnikov/blueprint/foundation/logger"
	"github.com/dmitryovchinnikov/blueprint/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log  *logger.Logger
	Auth *auth.Auth
	DB   *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	usmCore := usersummary.NewCore(usersummarydb.NewStore(cfg.Log, cfg.DB))

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	hdl := New(usmCore)
	app.Handle(http.MethodGet, version, "/usersummary", hdl.Query, authen, ruleAdmin)
}
