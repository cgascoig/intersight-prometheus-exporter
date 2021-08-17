package ismetrics

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type SourceConfig struct {
	KeyID       string `json:"key_id,omitempty"`
	KeyData     string `json:"key_data,omitempty"`
	KeyFilename string `json:"key_filename,omitempty"`
}

type SourceConfigList []SourceConfig

func GetSourceConfigListFromFile(filename string) (SourceConfigList, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading source configuration list file: %v", err)
	}

	return getSourceConfigList(b)
}

func GetSourceConfigListFromURL(url string) (SourceConfigList, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error requesting source configuration list file: %v", err)
	}

	b, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Error requesting source configuration list file: %v", err)
	}

	return getSourceConfigList(b)
}

func getSourceConfigList(data []byte) (SourceConfigList, error) {
	ret := SourceConfigList{}

	err := json.Unmarshal(data, &ret)
	if err != nil {
		return nil, fmt.Errorf("Error parsing source configuration list: %v", err)
	}

	return ret, nil
}

func (sc *SourceConfig) GetKeyData() string {
	if sc.KeyData != "" {
		return sc.KeyData
	}

	if sc.KeyFilename != "" {
		b, err := ioutil.ReadFile(sc.KeyFilename)
		if err != nil {
			return ""
		}

		return string(b)
	}

	return ""
}

func (sc *SourceConfig) String() string {
	return sc.KeyID
}
