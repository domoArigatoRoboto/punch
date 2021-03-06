// +build all integration

package cmdtest

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/require"
)

func TestAskForLogin(t *testing.T) {
	defer createConfig(t)()
	p := testcli.Command(exePath, "subdomain", "list", "--config", configPath)
	p.Run()

	require.Contains(t, p.Stderr(), "You need to login using `punch login` first.")
}

func TestLogin(t *testing.T) {
	defer createConfig(t)()
	p := testcli.Command(exePath, "login", "--config", configPath)
	p.Run()
	if !p.Failure() {
		t.Fatalf("Expected punch login to fail, but it succeed.")
	}

	if !p.StderrContains("required flag(s) \"password\", \"username\" not set") {
		t.Fatalf("Expected password and username to be required.")
	}
}

func TestLoginSetsTOML(t *testing.T) {
	defer createConfig(t)()
	p := testcli.Command(exePath, "login", "-u", "testuser@holepunch.io", "-p", "secret", "--config", configPath)
	p.Run()

	if !p.Success() {
		t.Fatalf("Expected punch login to succeed, but it failed.")
	}

	fmt.Println(p.Stdout())

	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		t.Fatal("/tmp/punch.toml not written")
	}

	require.Contains(t, string(dat), "apikey = \"eyJ0eXAiO")

}
func TestIncorrectLogin(t *testing.T) {
	defer createConfig(t)()
	p := testcli.Command(exePath, "login", "-u", "testuser@holepunch.io", "-p", "wrongpass", "--config", configPath)
	p.Run()

	if !p.Success() {

	} else {
		t.FailNow()
	}

}
