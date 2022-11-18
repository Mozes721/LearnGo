package contexts

import (
	"context"
	"fmt"
	"net/http"
)

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) CallAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
	}
	id, err := processResponse(resp.Body)
	return id, err
}
