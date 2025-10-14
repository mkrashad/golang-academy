package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	http *resty.Client
}

type PeopleResponse struct {
	Results []any `json:"results"`
}

func NewClient() *Client {
	return &Client{http: resty.New()}
}

func (c *Client) StarWarsCharacter(character string) {
	resp, err := c.http.R().
		SetQueryParam("search", character).
		Get("https://swapi.dev/api/people")

	if err != nil {
		log.Fatal(err)
	}

	var respBody PeopleResponse
	if err = json.Unmarshal(resp.Body(), &respBody); err != nil {
		log.Fatal(err)
	}

	if len(respBody.Results) < 1 {
		fmt.Println("No character found")
	} else {
		fmt.Println("Found character:", character, respBody.Results)
	}

	fmt.Println("Status Code:", resp.StatusCode())
}
