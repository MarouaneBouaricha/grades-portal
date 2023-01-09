package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		ServicesURL,
		bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service. Registry service responded with code %v", res.StatusCode)
	}
	return err
}

type providers struct {
	services map[ServiceName][]string
	mutex    *sync.RWMutex
}

func (p *providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

}

var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}
