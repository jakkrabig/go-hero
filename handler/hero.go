package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func GetHeros(c echo.Context) error {
	name := c.Param("name")
	var heros map[string]interface{}

	req, err := http.NewRequest("GET", "https://superheroapi.com/api/605647239925674/search/"+name, nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &heros)
	if err != nil {
		return err
	}

	return c.JSON(200, heros["results"])
}

func GetDetail(c echo.Context) error {
	id := c.Param("id")
	var heros map[string]interface{}

	req, err := http.NewRequest("GET", "https://superheroapi.com/api/605647239925674/"+id, nil)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &heros)
	if err != nil {
		return err
	}

	return c.JSON(200, heros)
}
