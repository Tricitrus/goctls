package generator

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	conf "github.com/Tricitrus/goctls/config"
	"github.com/Tricitrus/goctls/rpc/parser"
	"github.com/Tricitrus/goctls/util"
	"github.com/Tricitrus/goctls/util/pathx"
)

func (g *Generator) GenBaseDesc(ctx DirContext, _ parser.Proto, cfg *conf.Config, c *ZRpcContext) error {
	protoFilename := filepath.Join(ctx.GetMain().Filename, "desc", "base.proto")
	if err := pathx.MkdirIfNotExist(filepath.Join(ctx.GetMain().Filename, "desc")); err != nil {
		return err
	}

	err := util.With("t").Parse(rpcTemplateText).SaveTo(map[string]string{
		"package":     strings.ToLower(c.RpcName),
		"serviceName": strcase.ToCamel(c.RpcName),
	}, protoFilename, false)
	return err
}
