package fronius

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Client) GetRealtimeData(pv chan<- Inverter) error {
	rel := &url.URL{Path: "GetInverterRealtimeData.cgi"}
	url := baseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()

	q.Add("Scope", "System")
	req.URL.RawQuery = q.Encode()

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	var response Response[Inverter]
	err = json.NewDecoder(res.Body).Decode(&response)
	pv <- response.Body
	return err
}

func (i Inverter) MarshalJSON() ([]byte, error) {
	inverter := struct {
		Power        output `json:"power"`
		DailyEnergy  output `json:"daily_energy"`
		AnnualEnergy output `json:"annual_energy"`
		TotalEnergy  output `json:"total_energy"`
	}{
		Power:        i.Data.PAC,
		DailyEnergy:  i.Data.DAY_ENERGY,
		AnnualEnergy: i.Data.YEAR_ENERGY,
		TotalEnergy:  i.Data.TOTAL_ENERGY,
	}

	return json.Marshal(inverter)
}

func (o output) MarshalJSON() ([]byte, error) {
	output := struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	}{
		Value: o.Values.Sum(),
		Unit:  o.Unit,
	}

	return json.Marshal(output)
}

func (v values) Sum() int {
	sum := 0
	for _, t := range v {
		sum += t
	}

	return sum
}

type Inverter struct {
	Data struct {
		PAC          output
		DAY_ENERGY   output
		YEAR_ENERGY  output
		TOTAL_ENERGY output
	}
}

type output struct {
	Unit   string
	Values values
}

type values map[string]int
