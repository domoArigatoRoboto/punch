package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/rendon/testcli"
	"github.com/stretchr/testify/require"
)

func TestAskForLogin(t *testing.T) {
	defer CreateConfig(t)()
	p := testcli.Command("../punch", "subdomain", "list", "--config", CONFIG_PATH)
	p.Run()

	require.Contains(t, p.Stdout(), "You need to login using `punch login` first.")
}

func TestLogin(t *testing.T) {
	defer CreateConfig(t)()
	p := testcli.Command("../punch", "login", "--config", CONFIG_PATH)
	p.Run()
	if !p.Failure() {
		t.Fatalf("Expected punch login to fail, but it succeeed.")
	}

	if !p.StdoutContains("required flag(s) \"password\", \"username\" not set") {
		t.Fatalf("Expected password and username to be required.")
	}
}

func TestLoginSetsTOML(t *testing.T) {
	defer CreateConfig(t)()
	p := testcli.Command("../punch", "login", "-u", "testuser@holepunch.io", "-p", "secret", "--config", CONFIG_PATH)
	p.Run()

	if !p.Success() {
		t.Fatalf("Expected punch login to succeeed, but it failed.")
	}

	fmt.Println(p.Stdout())

	dat, err := ioutil.ReadFile(CONFIG_PATH)
	if err != nil {
		t.Fatal("/tmp/punch.toml not written")
	}

	require.Contains(t, string(dat), "apikey = \"eyJ0eXAiO")

}
func TestIncorrectLogin(t *testing.T) {
	defer CreateConfig(t)()
	p := testcli.Command("../punch", "login", "-u", "testuser@holepunch.io", "-p", "wrongpass", "--config", CONFIG_PATH)
	p.Run()

	if !p.Success() {

	} else {
		t.FailNow()
	}

}
