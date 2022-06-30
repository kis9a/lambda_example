package models

type Remo struct {
	RemoTable string `json:"remoTable"`
}

type RemoSensorValue struct {
	CreatedAt string  `json:"created_at,omitempty"`
	Val       float64 `json:"val,omitempty"`
}

type DeviceNewestEvents struct {
	Hu RemoSensorValue `json:"hu,omitempty"`
	Il RemoSensorValue `json:"il,omitempty"`
	Mo RemoSensorValue `json:"mo,omitempty"`
	Te RemoSensorValue `json:"te,omitempty"`
}

type RemoDevice struct {
	CreatedAt         string             `json:"created_at,omitempty"`
	FirmwareVersion   string             `json:"firmware_version,omitempty"`
	HumidityOffset    float64            `json:"humidity_offset,omitempty"`
	ID                string             `json:"id,omitempty"`
	MacAddress        string             `json:"mac_address,omitempty"`
	Name              string             `json:"name,omitempty"`
	NewestEvents      DeviceNewestEvents `json:"newest_events,omitempty"`
	SerialNumber      string             `json:"serial_number,omitempty"`
	TemperatureOffset float64            `json:"temperature_offset,omitempty"`
	UpdatedAt         string             `json:"updated_at,omitempty"`
}

func (r *Remo) SaveRemoA2DeviceEvents() {
}
