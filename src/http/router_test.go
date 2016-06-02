package http

import (
	"github.com/geodan/gost/src/configuration"
	"github.com/geodan/gost/src/database/postgis"
	"github.com/geodan/gost/src/mqtt"
	"github.com/geodan/gost/src/sensorthings/api"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test the router functionality
func TestCreateRouter(t *testing.T) {
	// arrange
	cfg := configuration.Config{}
	mqttServer := mqtt.CreateMQTTClient(configuration.MQTTConfig{})
	database := postgis.NewDatabase("", 123, "", "", "", "", false, 50, 100)
	a := api.NewAPI(database, cfg, mqttServer)

	// act
	router := CreateRouter(&a)

	// assert
	assert.NotNil(t, router, "Router should be created")
}
