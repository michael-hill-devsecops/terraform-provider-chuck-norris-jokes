package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() *schema.Provider {
            return &schema.Provider{
                Schema: map[string]*schema.Schema{},
                ResourcesMap: map[string]*schema.Resource{
                    "chucknorris_joke": resourceJoke(),
                },
            }
        },
    })
}

func resourceJoke() *schema.Resource {
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

func resourceJokeCreate(d *schema.ResourceData, m interface{}) error {
    // The Create function should ideally call Read to set the initial state
    return resourceJokeRead(d, m)
}

func resourceJokeRead(d *schema.ResourceData, m interface{}) error {
    joke, err := getJoke()
    if err != nil {
        return err
    }

    d.SetId("chucknorris_joke") // Consider using a unique identifier
    d.Set("joke", joke)
    return nil
}

func getJoke() (string, error) {
    resp, err := http.Get("https://api.chucknorris.io/jokes/random")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to get joke: %s", resp.Status)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]interface{}
    err = json.Unmarshal(body, &result)
    if err != nil {
        return "", err
    }

    joke, ok := result["value"].(string)
    if !ok {
        return "", fmt.Errorf("unexpected response format")
    }

    return joke, nil
}

func resourceJokeDelete(d *schema.ResourceData, m interface{}) error {
    // No-op: Nothing to delete
    d.SetId("") // Remove the ID to mark it as deleted
    return nil
}
