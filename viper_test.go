package main

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	var config *viper.Viper = viper.New()
	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "http://localhost:3306", config.GetString("database.host"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestYaml(t *testing.T) {
	config := viper.New()
	//config.SetConfigType("yaml")
	//config.SetConfigName("config")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("app.name"))
	assert.Equal(t, "http://localhost:3306", config.GetString("database.host"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
}

func TestEnvFile(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.Nil(t, err)
	assert.Equal(t, "golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, "http://localhost:3306", config.GetString("DATABASE_HOST"))
	assert.Equal(t, true, config.GetBool("DATABASE_SHOW_SQL"))
}
