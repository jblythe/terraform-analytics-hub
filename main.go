package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/jblythe/terraform-analytics-hub/analytics_hub"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: analytics_hub.Provider,
	})
}
