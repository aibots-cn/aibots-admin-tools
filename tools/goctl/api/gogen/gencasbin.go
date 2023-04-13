package gogen

import (
	_ "embed"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
)

//go:embed authortymiddleware.tpl
var authorityMiddlewareTemplate string

func genCasbin(dir string, cfg *config.Config, api *spec.ApiSpec) error {
	fileName, err := format.FileNamingFormat(cfg.NamingFormat, "authority_middleware.go")
	if err != nil {
		return err
	}

	err = genFile(fileGenConfig{
		dir:             dir,
		subdir:          middlewareDir,
		filename:        fileName,
		templateName:    "authorityTemplate",
		category:        category,
		templateFile:    authorityTemplateFile,
		builtinTemplate: authorityMiddlewareTemplate,
		data:            map[string]string{},
	})
	if err != nil {
		return err
	}

	return nil
}
