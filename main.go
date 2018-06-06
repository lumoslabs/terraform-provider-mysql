package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/lumoslabs/terraform-provider-mysql/mysql"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mysql.Provider})
}
