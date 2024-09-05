// resource_joke.go
//
// This file defines the Chuck Norris joke resource for the Terraform provider.
// It includes the schema for the resource and the CRUD (Create, Read, Delete)
// operations required by Terraform to manage this resource.
//
// The resource fetches a random Chuck Norris joke from an external API and
// makes it available as an output attribute in Terraform.

package resources

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/michael-hill-devsecops/terraform-provider-chucknorris/handlers"
)

// ResourceJoke returns a schema.Resource for the Chuck Norris joke
func ResourceJoke() *schema.Resource {
    return &schema.Resource{
        Create: resourceJokeCreate,
        Read:   resourceJokeRead,
        Delete: resourceJokeDelete,
        Schema: map[string]*schema.Schema{
            "joke": {
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}

// resourceJokeCreate is called when a new resource is created
func resourceJokeCreate(d *schema.ResourceData, m interface{}) error {
    // Fetch the joke and set the initial state
    return resourceJokeRead(d, m)
}

// resourceJokeRead retrieves the current state of the resource
func resourceJokeRead(d *schema.ResourceData, m interface{}) error {
    joke, err := handlers.GetJoke()
    if err != nil {
        return err
    }

    // Set the ID to a fixed value (optional: could use a unique value)
    d.SetId("chucknorris_joke")
    d.Set("joke", joke)
    return nil
}

// resourceJokeDelete handles deletion of the resource
func resourceJokeDelete(d *schema.ResourceData, m interface{}) error {
    // No-op: No state to delete
    d.SetId("") // Remove the ID to mark the resource as deleted
    return nil
}
