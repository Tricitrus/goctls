package config

import (
    {{if .useCasbin}}"github.com/Tricitrus/tricitrus-admin-common/plugins/casbin"
    "github.com/Tricitrus/tricitrus-admin-common/config"{{else}}{{if .useEnt}}"github.com/Tricitrus/tricitrus-admin-common/config"{{end}}{{end}}{{if .useI18n}}
	"github.com/Tricitrus/tricitrus-admin-common/i18n"
{{end}}
	"github.com/zeromicro/go-zero/rest"{{if .useCoreRpc}}
	"github.com/zeromicro/go-zero/zrpc"{{end}}
)

type Config struct {
	rest.RestConf
	Auth         rest.AuthConf
	CROSConf     config.CROSConf
	{{if .useCasbin}}CasbinDatabaseConf config.DatabaseConf
    RedisConf    config.RedisConf
	CasbinConf   casbin.CasbinConf{{end}}{{if .useEnt}}
	DatabaseConf config.DatabaseConf{{end}}{{if .useCoreRpc}}
	CoreRpc      zrpc.RpcClientConf{{end}}{{if .useI18n}}
	I18nConf     i18n.Conf{{end}}
	{{.jwtTrans}}
}
