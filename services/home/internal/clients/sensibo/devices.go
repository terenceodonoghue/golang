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
	var response Response[Device]
	err = json.NewDecoder(res.Body).Decode(&response)
	return response.Result, err
}

func (d *Device) MarshalJSON() ([]byte, error) {
	device := struct {
		Id                 string      `json:"id,omitempty"`
		Room               string      `json:"room,omitempty"`
		Mode               string      `json:"mode,omitempty"`
		FanLevel           string      `json:"fan_level,omitempty"`
		Running            bool        `json:"running,omitempty"`
		Humidity           float32     `json:"humidity,omitempty"`
		CurrentTemperature Temperature `json:"current_temperature,omitempty"`
		TargetTemperature  Temperature `json:"target_temperature,omitempty"`
	}{
		Id:       d.Id,
		Room:     d.Room.Name,
		Mode:     d.AcState.Mode,
		FanLevel: d.AcState.FanLevel,
		Running:  d.AcState.On,
		Humidity: d.Measurements.Humidity,
		CurrentTemperature: Temperature{
			Value: d.Measurements.Temperature,
			Unit:  d.TemperatureUnit,
		},
		TargetTemperature: Temperature{
			Value: d.AcState.TargetTemperature,
			Unit:  d.AcState.TemperatureUnit,
		},
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

type Temperature struct {
	Value float32 `json:"value"`
	Unit  string  `json:"unit"`
}

type State struct {
	On                bool
	Mode              string
	FanLevel          string
	TargetTemperature float32
	TemperatureUnit   string
}
