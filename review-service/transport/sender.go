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
type BookInfo struct {
	Author *string
	Title  *string
	Error  *error
}

type UserInfo struct {
	Name  *string
	Error *error
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

func FetchBookInfo(ctx context.Context, cfg bootstrap.Config, bookId uint, ch chan BookInfo) error {
	url := cfg.BookServiceUrl

	resp, err := fetchInfo(ctx, url, bookId)
	if err != nil {
		return err
	}
	bookInfo, err := DeserializeBookInfo(resp)
	if err != nil {
		return err
	}
	author := bookInfo.GetAuthor()
	title := bookInfo.GetTitle()

	ch <- BookInfo{
		Author: &author,
		Title:  &title,
		Error:  nil,
	}
	return nil
}

func FetchUserInfo(ctx context.Context, cfg bootstrap.Config, userId uint, ch chan UserInfo) error {
	url := cfg.UserServiceUrl

	resp, err := fetchInfo(ctx, url, userId)
	if err != nil {
		return err
		// ch <- UserInfo{nil, &err}
	}
	userInfo, err := DeserializeUserInfo(resp)
	if err != nil {
		return err
		// ch <- UserInfo{nil, &err}
	}
	name := userInfo.GetName()

	ch <- UserInfo{
		Name:  &name,
		Error: nil,
	}
	return nil
}
