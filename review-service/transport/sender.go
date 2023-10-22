package transport

import (
	"context"
	"log"
	"net/http"
	"review-service/bootstrap"
	"strconv"
)

type ServiceInfo struct {
	Client  *http.Client
	Context context.Context
	URL     string
}

func NewServiceInfo(context context.Context, ServiceURL string, EntityID uint) *ServiceInfo {

	strId := strconv.FormatUint(uint64(EntityID), 10)

	return &ServiceInfo{
		Client:  &http.Client{},
		Context: context,
		URL:     ServiceURL + strId,
	}
}

func fetchInfo(ctx context.Context, serviceUrl string, entityId uint) (*http.Response, error) {

	serviceInfo := NewServiceInfo(ctx, serviceUrl, entityId)

	req, err := http.NewRequestWithContext(ctx, "GET", serviceInfo.URL, nil)
	log.Printf("Sending HTTP request to: %v", serviceInfo.URL)

	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/x-protobuf")

	resp, err := serviceInfo.Client.Do(req)
	if err != nil {
		log.Printf("Error making HTTP request: %v", err)
		return nil, err
	}

	return resp, nil
}

func FetchBookInfo(ctx context.Context, cfg bootstrap.Config, bookId uint) (*http.Response, error) {
	url := cfg.BookServiceUrl
	return fetchInfo(ctx, url, bookId)
}

func FetchUserInfo(ctx context.Context, cfg bootstrap.Config, userId uint) (*http.Response, error) {
	url := cfg.UserServiceUrl
	return fetchInfo(ctx, url, userId)
}
