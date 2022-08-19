package outbox

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DebeziumSettings struct {
	Url               string
	ConnectorFilePath string
}

type RegisterDebezium struct {
	Settings DebeziumSettings
}

func (d *RegisterDebezium) CheckConnector(connectorName string) (bool, error) {
	response, err := http.Get(fmt.Sprintf("%s/%s", d.Settings.Url, connectorName))
	defer response.Body.Close()

	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return true, nil
	}
	return false, errors.New("failed to connect")
}
func (d *RegisterDebezium) RegisterConnector() *http.Response {
	plan, _ := ioutil.ReadFile(d.Settings.ConnectorFilePath)
	response, err := http.Post(d.Settings.Url, "application/json", bytes.NewBuffer(plan))

	if err != nil {
		panic(err)
	}

	return response
}
func NewRegisterDebezium(settings DebeziumSettings) *RegisterDebezium {
	return &RegisterDebezium{
		Settings: settings,
	}
}
