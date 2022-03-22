package api

import "testing"

func TestGetDataFromCity(t *testing.T) {
	t.Log("Init get data test from city")

	data, err := OpenWeatherGetDataFromCity("cuiaba")

	if err != nil {
		t.Fatalf("error getting data: %v", data)
	}

	if data.Name != "Cuiabá" {
		t.Fatalf("expect receive Cuiabá inside %s", data.Name)
	}
}
