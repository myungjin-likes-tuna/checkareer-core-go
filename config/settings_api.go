package config

import "fmt"

// BindAddress 서버 주소
func (s Settings) BindAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}
