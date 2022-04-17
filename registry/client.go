package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "applicaiton/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. registry service responded with code %v", res.StatusCode)
	}
	return nil
}

func DerigesterService(serviceUrl string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(serviceUrl)))
	req.Header.Add("Content-Type", "text/plain")
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to de-register service. registry service responded with code %v", res.StatusCode)
	}
	return err
}
