package homegrp

import (
	"net/http"

	"github.com/dmitryovchinnikov/blueprint/business/core/event"
	"github.com/dmitryovchinnikov/blueprint/business/core/home"
	"github.com/dmitryovchinnikov/blueprint/business/core/home/stores/homedb"
	"github.com/dmitryovchinnikov/blueprint/business/core/user"
	"github.com/dmitryovchinnikov/blueprint/business/core/user/stores/usercache"
	"github.com/dmitryovchinnikov/blueprint/business/core/user/stores/userdb"
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

	envCore := event.NewCore(cfg.Log)
	usrCore := user.NewCore(cfg.Log, envCore, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB)))
	hmeCore := home.NewCore(cfg.Log, envCore, usrCore, homedb.NewStore(cfg.Log, cfg.DB))

	authen := mid.Authenticate(cfg.Auth)
	tran := mid.ExecuteInTransation(cfg.Log, db.NewBeginner(cfg.DB))

	hdl := New(hmeCore)
	app.Handle(http.MethodGet, version, "/homes", hdl.Query, authen)
	app.Handle(http.MethodGet, version, "/homes/:home_id", hdl.QueryByID, authen)
	app.Handle(http.MethodPost, version, "/homes", hdl.Create, authen)
	app.Handle(http.MethodPut, version, "/homes/:home_id", hdl.Update, authen, tran)
	app.Handle(http.MethodDelete, version, "/homes/:home_id", hdl.Delete, authen, tran)
}
