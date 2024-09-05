// main.go
//
// This file serves as the entry point for the Terraform provider plugin. It
// initializes and serves the custom Terraform provider by defining the
// provider function and mapping the custom resource.
//
// The provider uses the schema from the "resources" package to define the
// available resources, in this case, a Chuck Norris joke resource.

package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
    "github.com/michael-hill-devsecops/terraform-provider-chucknorris/resources"
)

// main function to serve the Terraform provider plugin
func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() *schema.Provider {
            return &schema.Provider{
                Schema:       map[string]*schema.Schema{},
                ResourcesMap: map[string]*schema.Resource{"chucknorris_joke": resources.ResourceJoke()},
            }
        },
    })
}
