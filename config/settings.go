package config

import (
	"encoding/json"
	"log"

	env "github.com/Netflix/go-env"
	"go.uber.org/fx"
)

// Settings 설정
type Settings struct {
	Neo4j struct {
		URI      string `env:"NEO4J_URI" json:"uri"`
		Username string `env:"NEO4J_USERNAME" json:"username"`
		Password string `env:"NEO4J_PASSWORD" json:"password"`
	} `json:"neo4j"`
	Port   int        `json:"port"  env:"PORT,default=7279"`
	CI     bool       `json:"ci"    env:"CI"`
	Extras env.EnvSet `json:"-"`
}

// NewSettings 설정 값 생성
func NewSettings() *Settings {
	var settings Settings
	extras, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Fatal(err)
	}
	settings.Extras = extras
	return &settings
}

// JSON 설정 값 출력
func JSON() string {
	settings := NewSettings()
	jsonBytes, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonBytes)
}

// Modules config 모듈
var Modules = fx.Options(fx.Provide(NewSettings))
