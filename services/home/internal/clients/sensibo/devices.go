package sensibo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetDevices(ac chan<- []Device) error {
	rel := &url.URL{Path: "users/me/pods"}
	url := baseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	f := []string{"id", "acState", "measurements", "room", "temperatureUnit"}

	q.Add("apiKey", c.apiKey)
	q.Add("fields", strings.Join(f, ","))
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var response Response[Device]
	err = json.NewDecoder(res.Body).Decode(&response)
	ac <- response.Result
	return err
}

func (d *Device) MarshalJSON() ([]byte, error) {
	device := struct {
		Id                 string      `json:"id"`
		Room               string      `json:"room"`
		Mode               string      `json:"mode"`
		FanLevel           string      `json:"fan_level"`
		IsRunning          bool        `json:"is_running"`
		Humidity           float32     `json:"humidity"`
		CurrentTemperature temperature `json:"current_temperature"`
		TargetTemperature  temperature `json:"target_temperature"`
	}{
		Id:        d.Id,
		Room:      d.Room.Name,
		Mode:      d.AcState.Mode,
		FanLevel:  d.AcState.FanLevel,
		IsRunning: d.AcState.On,
		Humidity:  d.Measurements.Humidity,
		CurrentTemperature: temperature{
			Value: d.Measurements.Temperature,
			Unit:  d.TemperatureUnit,
		},
		TargetTemperature: temperature{
			Value: d.AcState.TargetTemperature,
			Unit:  d.AcState.TemperatureUnit,
		},
	}

	return json.Marshal(device)
}

type Device struct {
	Id              string
	Room            room
	AcState         state
	Measurements    measurement
	TemperatureUnit string
}

type measurement struct {
	Temperature float32
	Humidity    float32
}

type room struct {
	Name string
}

type state struct {
	On                bool
	Mode              string
	FanLevel          string
	TargetTemperature float32
	TemperatureUnit   string
}

type temperature struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}
