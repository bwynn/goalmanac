package models

type Observation struct {
	ObservationEpoch string    `json:"observation_epoch"`
	StationID        string    `json:"station_id"`
	TempC            float64   `json:"temp_c"`
	WindDegrees      int       `json:"wind_degrees"`
	WindKPH          float64   `json:"wind_kph"`
	WindGustKPH      string    `json:"wind_gust_kph"`
	PressureMB       string    `json:"pressure_mb"`
	PressureIN       string    `json:"pressure_in"`
	PressureTrend    string    `json:"pressure_trend"`
	DewpointC        int       `json:"dewpoint_c"`
	WindchillC       string    `json:"windchill_c"`
	Precip1HRMetric  string    `json:"precip_1hr_metric"`
	DisplayLocation  *Location `json:"display_location"`
}

type Location struct {
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
	Elevation string `json:"elevation"`
}

type CurrentObservations struct {
	CurrentObservation *Observation `json:"current_observation"`
}
