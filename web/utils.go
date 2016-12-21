package web

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"syscall"

	"github.com/spf13/viper"
)

// HostURL get host url
func HostURL() string {
	if IsProduction() {
		return fmt.Sprintf("https://%s", viper.GetString("server.name"))
	}
	return fmt.Sprintf("http://localhost:%d", viper.GetInt("server.port"))
}

// IsProduction production mode ?
func IsProduction() bool {
	return viper.GetString("env") == "production"
}

//Shell exec shell command
func Shell(cmd string, args ...string) error {
	bin, err := exec.LookPath(cmd)
	if err != nil {
		return err
	}
	return syscall.Exec(bin, append([]string{cmd}, args...), os.Environ())
}

//RandomStr randome string
func RandomStr(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	buf := make([]rune, n)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	return string(buf)
}
