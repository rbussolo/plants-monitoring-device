package device

import "testing"

func TestGetData(t *testing.T) {
	t.Log("Init get data from device")

	d := NewDevice("TESTE_DEVICE")

	// Get data from device
	data, err := d.GetData()

	if err != nil {
		t.Fatalf("error getting data from api %v", err)
	}

	if data.timestamp.IsZero() {
		t.Fatal("not found data from device")
	}
}

func TestSendData(t *testing.T) {
	t.Log("Init send data to server")

	d := NewDevice("TESTE_DEVICE")

	data, err := d.GetData()
	if err != nil {
		t.Fatalf("error getting data from api %v", err)
	}

	// Put the data at device
	d.lastData = data

	// Send data to the server
	success, err := d.SendData()
	if !success {
		t.Fatalf("error sending data: %v", err)
	}
}
