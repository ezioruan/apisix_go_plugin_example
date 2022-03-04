package main

import (
	"log"

	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"github.com/apache/apisix-go-plugin-runner/pkg/runner"
	"github.com/ezioruan/apisix_go_plugin_example/plugins"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := runner.RunnerConfig{}
	cfg.LogLevel = zapcore.DebugLevel
	err := plugin.RegisterPlugin(&plugins.BasicAuthPlugin{})
	if err != nil {
		log.Fatalf("failed to register plugin BasicAuthPlugin: %s", err)
	}
	runner.Run(cfg)
}
