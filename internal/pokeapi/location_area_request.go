package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreaResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint
	if pageURL != nil {
		fullUrl = *pageURL
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResp{}, err
	}

	locationAreaResp := LocationAreaResp{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil {
		return LocationAreaResp{}, err
	}

	return locationAreaResp, nil
}

func (c *Client) FetchLocationAreaDetail(nameOrID string) (LocationAreaDetail, error) {
	var detail LocationAreaDetail
	url := fmt.Sprintf("%s/location-area/%s", baseURL, nameOrID)

	resp, err := http.Get(url)
	if err != nil {
		return detail, fmt.Errorf("error fetching location area detail: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return detail, fmt.Errorf("PokeAPI returned non-OK status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&detail)
	if err != nil {
		return detail, fmt.Errorf("error decoding response: %v", err)
	}

	return detail, nil
}
