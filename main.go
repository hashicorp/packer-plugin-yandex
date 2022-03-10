package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	"github.com/hashicorp/packer-plugin-yandex/builder/yandex"
	yandexexport "github.com/hashicorp/packer-plugin-yandex/post-processor/yandex-export"
	yandeximport "github.com/hashicorp/packer-plugin-yandex/post-processor/yandex-import"
	"github.com/hashicorp/packer-plugin-yandex/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(yandex.Builder))
	pps.RegisterPostProcessor("import", new(yandeximport.PostProcessor))
	pps.RegisterPostProcessor("export", new(yandexexport.PostProcessor))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
