package config

import (
	"fmt"
	"os"
	"strings"
)

var (
	PORT                           string
	GOOGLE_APPLICATION_CREDENTIALS string
	LINE_CHANNEL_ACCESS_TOKEN      string
	MODE                           Mode
	HOST                           HostName
)

type (
	Mode     string
	HostName string
)

var (
	Production Mode = "production"
	Developing Mode = "developing"
)

func init() {
	var err error

	PORT, err = getPORT()
	if err != nil {
		panic(err)
	}

	GOOGLE_APPLICATION_CREDENTIALS, err = getGOOGLE_APPLICATION_CREDENTIALS()
	if err != nil {
		panic(err)
	}

	LINE_CHANNEL_ACCESS_TOKEN, err = getLINE_CHANNEL_ACCESS_TOKEN()
	if err != nil {
		panic(err)
	}

	MODE, err = getMODE()
	if err != nil {
		panic(err)
	}

	HOST, err = getHOST()
	if err != nil {
		panic(err)
	}

}

func getPORT() (string, error) {
	key := "PORT"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getGOOGLE_APPLICATION_CREDENTIALS() (string, error) {
	key := "GOOGLE_APPLICATION_CREDENTIALS"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getLINE_CHANNEL_ACCESS_TOKEN() (string, error) {
	key := "LINE_CHANNEL_ACCESS_TOKEN"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return e, nil
}

func getMODE() (Mode, error) {
	var m Mode
	if s := os.Getenv("MODE"); s == "production" {
		m = Production
	} else {
		m = Developing
	}
	return m, nil
}

func getHOST() (HostName, error) {
	key := "HOST"
	e := os.Getenv(key)
	if e == "" {
		return "", fmt.Errorf("the environment variable %s must be filled", key)
	}
	return HostName(e), nil
}

func (h HostName) File(path string) string {
	host := strings.Split(path, ":")
	fmt.Println(host)
	if MODE == Production {
		return fmt.Sprintf("https://%s/%s", h, path)
	} else {
		return fmt.Sprintf("http://%s/%s", h, path)
	}
}
