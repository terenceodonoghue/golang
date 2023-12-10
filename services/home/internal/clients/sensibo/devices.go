package sensibo

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetDevices() ([]Device, error) {
	rel := &url.URL{Path: "users/me/pods"}
	url := baseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	f := []string{"id", "acState", "measurements", "room", "temperatureUnit"}

	q.Add("apiKey", c.apiKey)
	q.Add("fields", strings.Join(f, ","))
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var devices Response[Device]
	err = json.NewDecoder(res.Body).Decode(&devices)
	return devices.Result, err
}

func (d *Device) MarshalJSON() ([]byte, error) {
	device := struct {
		Id                 string  `json:"id,omitempty"`
		Room               string  `json:"room,omitempty"`
		Mode               string  `json:"mode,omitempty"`
		FanLevel           string  `json:"fan_level,omitempty"`
		TemperatureUnit    string  `json:"temperature_unit,omitempty"`
		CurrentTemperature float32 `json:"current_temperature,omitempty"`
		TargetTemperature  float32 `json:"target_temperature,omitempty"`
		Humidity           float32 `json:"humidity,omitempty"`
		Running            bool    `json:"running,omitempty"`
	}{
		Id:                 d.Id,
		Room:               d.Room.Name,
		Mode:               d.AcState.Mode,
		FanLevel:           d.AcState.FanLevel,
		TemperatureUnit:    d.TemperatureUnit,
		CurrentTemperature: d.Measurements.Temperature,
		TargetTemperature:  d.AcState.TargetTemperature,
		Humidity:           d.Measurements.Humidity,
		Running:            d.AcState.On,
	}

	return json.Marshal(device)
}

type Device struct {
	Id              string
	Room            Room
	AcState         State
	Measurements    Measurement
	TemperatureUnit string
}

type Measurement struct {
	Temperature float32
	Humidity    float32
}

type Room struct {
	Name string
}

type State struct {
	On                bool
	Mode              string
	FanLevel          string
	TargetTemperature float32
}
