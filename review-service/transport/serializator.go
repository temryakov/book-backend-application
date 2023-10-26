package transport

import (
	"fmt"
	"io"
	"log"
	"net/http"
	rp "review-service/proto"

	"google.golang.org/protobuf/proto"
)

func NewProtoDeserialization(resp *http.Response, message proto.Message) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if err := proto.Unmarshal(body, message); err != nil {
		return fmt.Errorf("error unmarshalling protobuf: %w", err)
	}

	log.Printf("Deserialized message: %+v", message)
	return nil
}

func DeserializeBookInfo(resp *http.Response) (*rp.GetBookResponse, error) {
	bookInfo := &rp.GetBookResponse{}
	if err := NewProtoDeserialization(resp, bookInfo); err != nil {
		return nil, err
	}
	return bookInfo, nil
}

func DeserializeUserInfo(resp *http.Response) (*rp.GetUserResponse, error) {
	userInfo := &rp.GetUserResponse{}
	if err := NewProtoDeserialization(resp, userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}
