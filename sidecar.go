package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var customPath string = ""

var cache map[string]map[string]string = make(map[string]map[string]string, 0)

func getPath() string {
	if len(customPath) != 0 {
		return customPath
	} else if os.Getenv("ENVIRONMENT") == "production" {
		return "/vault/secrets"
	} else {
		return "vault/secrets"
	}
}

func getFile(fileName string) ([]byte, error) {
	path := getPath()
	file, err := os.Open(filepath.Join(path, fileName))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, _ := ioutil.ReadAll(file)

	return bytes, nil
}

// Configure sidecar with a custom vault secrets path
// ex: /vault/secrets
func Configure(_path string) {
	customPath = _path
}

// Get secrets map from file name
// Secrets are cached after the first pull
func GetSecrets(fileName string) (map[string]string, error) {

	cacheVal := cache[fileName]

	if cacheVal != nil {
		return cacheVal, nil
	}

	bytes, err := getFile(fileName)
	if err != nil {
		return nil, err
	}

	var secrets map[string]string

	json.Unmarshal(bytes, &secrets)

	cache[fileName] = secrets

	return secrets, nil
}

// get database creds from file name
// returns (*DatabaseCreds, error)
func GetDatabaseCreds(fileName string) (*DatabaseCreds, error) {
	bytes, err := getFile(fileName)
	if err != nil {
		return nil, err
	}

	var creds DatabaseCreds

	json.Unmarshal(bytes, &creds)

	if len(creds.Username) == 0 || len(creds.Password) == 0 {
		return nil, errors.New("missing database username or password")
	}

	return &creds, nil
}

type DatabaseCreds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
