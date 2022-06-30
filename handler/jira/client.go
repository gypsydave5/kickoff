package jira

import (
	"encoding/json"
	"net/http"
)

type Fields struct {
	Summary string `json:"summary"`
	Status  `json:"status"`
}

type Status struct {
	CurrentStatus string `json:"name"`
}

type Story struct {
	Fields `json:"fields"`
	ID     string `json:"key"`
	Link   string
}

type Client struct {
	http *http.Client
}

func NewClient(http *http.Client) *Client {
	return &Client{http: http}
}

func (c Client) GetStory(storyID string) (Story, error) {
	var s Story
	req, _ := http.NewRequest(http.MethodGet, "https://saltpayco.atlassian.net/rest/api/3/issue/"+storyID, nil)
	res, err := c.http.Do(req)
	if err != nil {
		return s, err
	}

	s.Link = "https://saltpayco.atlassian.net/browse/" + storyID

	err = json.NewDecoder(res.Body).Decode(&s)
	return s, err
}
