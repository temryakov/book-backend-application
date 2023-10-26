package transport

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	rp "review-service/proto"

	"google.golang.org/protobuf/proto"
)

func DeserializeBookInfo(resp *http.Response) (*rp.GetBookResponse, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	bookInfo := &rp.GetBookResponse{}

	err = proto.Unmarshal(body, bookInfo)
	if err != nil {
		fmt.Println("Error unmarshalling protobuf:", err)
		return nil, err
	}
	log.Print(bookInfo)

	return bookInfo, nil
}

func DeserializeUserInfo(resp *http.Response) (*rp.GetUserResponse, error) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	userInfo := &rp.GetUserResponse{}

	err = proto.Unmarshal(body, userInfo)
	if err != nil {
		fmt.Println("Error unmarshalling protobuf:", err)
		return nil, err
	}
	log.Print(userInfo)

	return userInfo, nil
}
