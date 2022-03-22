package device

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	o "device/src/api"
)

var cities []string = []string{
	"cuiaba",
	"varzea grande",
	"sao paulo",
	"rio de janeiro",
	"curitiba",
	"cascavel",
	"rondonopolis",
	"sinop",
	"campo verde",
	"campo grande",
}

type User struct {
	Email     string
	ApiKey    string
	AuthToken string
}

type DeviceResquest struct {
	DeviceId     string  `json:"device"`
	Timestamp    int64   `json:"timestamp"`
	SoilMoisture int     `json:"soilMoisture"`
	Temperature  float32 `json:"temperature"`
	Humidity     int     `json:"humidity"`
}

type AuthRequest struct {
	ApiKey string `json:"api_key"`
	Email  string `json:"email"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type DeviceData struct {
	timestamp           time.Time
	soilMoisture        int
	externalTemperature float32
	externalHumidity    int
}

type Device struct {
	Id       string
	Interval int
	Api      string
	User     User
	City     string
	active   bool
	lastData DeviceData
}

func NewDevice(id string) *Device {
	return &Device{
		Id:       id,
		Interval: 10,
		Api:      "localhost:9090",
		User: User{
			Email:  "rbussolo91@gmail.com",
			ApiKey: "26a9ffa4-c373-4e2e-84a4-24561db63da5",
		},
		City: cities[rand.Intn(10)],
	}
}

func (d Device) Start() {
	var err error

	// Auth at server
	token, err := d.Auth()
	if err != nil {
		log.Fatalf("error auth: %v", err)
	}

	d.User.AuthToken = token

	initialHour := 0

	// Define this device is Active
	d.active = true

	// Define interval of send data
	intervalDuration := time.Duration(d.Interval) * time.Second

	// Loop each interval
	for range time.Tick(intervalDuration) {
		// Stopped device
		if !d.active {
			return
		}

		// Get data to send
		d.lastData, err = d.GetData()
		if err != nil {
			log.Fatalf("error getting data: %v", err)
		}

		// Add a hour on each data submission
		d.lastData.timestamp = d.lastData.timestamp.Add(time.Duration(initialHour) * time.Hour)

		// Send data to server
		_, err = d.SendData()
		if err != nil {
			log.Fatalf("error sending data: %v", err)
		}

		initialHour += 1
	}
}

func (d Device) Auth() (string, error) {
	var token string
	var authResponse AuthResponse

	ar := AuthRequest{
		ApiKey: d.User.ApiKey,
		Email:  d.User.Email,
	}

	jsonBody, err := json.Marshal(ar)

	if err != nil {
		return token, err
	}

	bytesBody := bytes.NewBuffer(jsonBody)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, d.Api+"/api/auth", bytesBody)
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)

	if err != nil {
		return token, err
	}

	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		return token, err
	}

	token = authResponse.Token

	return token, nil
}

func (d Device) Stop() {
	d.active = false
}

func (d Device) GetData() (DeviceData, error) {
	deviceData := DeviceData{}

	data, err := o.OpenWeatherGetDataFromCity(d.City)
	if err != nil {
		return deviceData, err
	}

	deviceData.timestamp = time.Unix(data.Dt, 0)
	deviceData.soilMoisture = data.Main.Pressure
	deviceData.externalHumidity = data.Main.Humidity
	deviceData.externalTemperature = data.Main.Temp

	return deviceData, err
}

func (d Device) SendData() (bool, error) {
	dj := DeviceResquest{
		DeviceId:     d.Id,
		Timestamp:    int64(time.Nanosecond) * d.lastData.timestamp.UnixNano() / int64(time.Millisecond),
		SoilMoisture: d.lastData.soilMoisture,
		Temperature:  d.lastData.externalTemperature,
		Humidity:     d.lastData.externalHumidity,
	}

	jsonBody, err := json.Marshal(dj)

	if err != nil {
		return false, err
	}

	bytesBody := bytes.NewBuffer(jsonBody)

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, d.Api+"/api/device/info", bytesBody)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", d.User.AuthToken)
	_, err = client.Do(req)

	if err != nil {
		return false, err
	}

	return true, nil
}
