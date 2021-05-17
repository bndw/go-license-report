package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const cacheDir = "go-license-report.cache"

func cacheGet(url string) (*license, bool) {
	key := cacheKey(url)
	fp := path.Join(cacheDir, key+".json")
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		Log.Printf("failed to read cache file: %w", err)
		return nil, false
	}

	var lic license
	if err := json.Unmarshal(b, &lic); err != nil {
		Log.Printf("failed to decode file %s: %w", fp, err)
		return nil, false
	}
	return &lic, true
}

func cacheSet(url string, lic *license) error {
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		os.Mkdir(cacheDir, os.ModePerm)
	}

	data, err := json.Marshal(lic)
	if err != nil {
		return err
	}

	key := cacheKey(url)
	fp := path.Join(cacheDir, key+".json")
	return ioutil.WriteFile(fp, data, 0644)
}

func cacheKey(v string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(v)))
}
