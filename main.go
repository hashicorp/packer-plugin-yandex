package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/hashicorp/packer-plugin-sdk/version"

	"github.com/hashicorp/packer-plugin-yandex/builder/yandex"
	yandexexport "github.com/hashicorp/packer-plugin-yandex/post-processor/yandex-export"
	yandeximport "github.com/hashicorp/packer-plugin-yandex/post-processor/yandex-import"
)

var (
	// Version is the main version number that is being run at the moment.
	Version = "0.0.1"

	// VersionPrerelease is A pre-release marker for the Version. If this is ""
	// (empty string) then it means that it is a final release. Otherwise, this
	// is a pre-release such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"

	// PluginVersion is used by the plugin set to allow Packer to recognize
	// what version this plugin is.
	PluginVersion = version.InitializePluginVersion(Version, VersionPrerelease)
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(yandex.Builder))
	pps.RegisterPostProcessor("import", new(yandeximport.PostProcessor))
	pps.RegisterPostProcessor("export", new(yandexexport.PostProcessor))
	pps.SetVersion(PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
