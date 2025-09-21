package generator

import (
	_ "embed"
	"path/filepath"

	"github.com/iancoleman/strcase"

	conf "github.com/Tricitrus/goctls/config"
	"github.com/Tricitrus/goctls/rpc/parser"
	"github.com/Tricitrus/goctls/util"
	"github.com/Tricitrus/goctls/util/format"
	"github.com/Tricitrus/goctls/util/pathx"
)

//go:embed enttx.tpl
var entTxTemplateText string

func (g *Generator) GenEntTx(ctx DirContext, _ parser.Proto, cfg *conf.Config, c *ZRpcContext) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, "ent_tx.go")
	if err != nil {
		return err
	}

	handlerFilename := filepath.Join(ctx.GetInternal().Filename, "utils/entx", filename)
	if err := pathx.MkdirIfNotExist(filepath.Join(ctx.GetInternal().Filename, "utils/entx")); err != nil {
		return err
	}

	err = util.With("entTx").Parse(entTxTemplateText).SaveTo(map[string]string{
		"package":     ctx.GetMain().Package,
		"serviceName": strcase.ToCamel(c.RpcName),
	}, handlerFilename, false)
	return err
}
