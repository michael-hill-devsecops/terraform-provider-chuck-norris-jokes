// joke_api.go
//
// This file contains the code to interact with the Chuck Norris API.
// It includes the GetJoke function, which fetches a random joke from
// the API and returns it. The function handles HTTP requests and JSON
// parsing to extract the joke text from the API response.

package handlers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// GetJoke fetches a random Chuck Norris joke from the API
func GetJoke() (string, error) {
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
