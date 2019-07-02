package source

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/qinhan-shu/consul/module"
)

// Github describes github repo of gmdata
type Github struct {
	user    string
	token   string
	repoURL string
}

// NewConfigSource : create config source
func NewConfigSource() (module.ConfigSource, error) {
	require := []string{
		"GITHUB_URL",
		"GITHUB_USERNAME",
		"GITHUB_TOKEN",
	}

	conf := make(map[string]string)
	for _, key := range require {
		value, isExist := os.LookupEnv(key)
		if !isExist {
			return nil, fmt.Errorf(`Environment "%s" must be set`, key)
		}
		conf[key] = value
	}

	return &Github{
		user:    conf["GITHUB_USERNAME"],
		token:   conf["GITHUB_TOKEN"],
		repoURL: conf["GITHUB_URL"],
	}, nil
}

// Fetch is to get details from github
func (g *Github) fetch(fileName string) ([]byte, error) {
	fileURL := fmt.Sprintf("%s/%s", g.repoURL, fileName)
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, err := http.NewRequest(http.MethodGet, fileURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github.VERSION.raw")
	req.SetBasicAuth(g.user, g.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if len(body) == 0 {
		return nil, module.ErrInvalidConfigJSON
	}

	if !json.Valid(body) {
		return nil, module.ErrInvalidConfigJSON
	}
	return body, nil
}
