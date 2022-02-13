package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/jmcvetta/napping"
	collectionpb "github.com/thanhlam/home-collect-data-svc/collectionpb"
	"github.com/thanhlam/home-collect-data-svc/service"
	"google.golang.org/grpc"
)

type server struct{}

//Check Token
func (*server) CheckToken(ctx context.Context, req *collectionpb.TokenRequest) (*collectionpb.TokenResponse, error) {
	log.Println("Server check token")
	fmt.Println("================")
	fmt.Println(req.GetTokenReq())
	fmt.Println("================")
	token := req.GetTokenReq()
	fmt.Println(token)
	status := CheckToken(token)
	/*if req.GetTokenReq() == "ThanhLam1" {
		resp := &collectionpb.TokenResponse{
			TokenRes: "YES",
		}
		return resp, nil
	} else {
		resp := &collectionpb.TokenResponse{
			TokenRes: "NO",
		}
		return resp, nil
	}*/
	if status == "ACTIVE" {
		resp := &collectionpb.TokenResponse{
			TokenRes: "YES",
		}
		return resp, nil
	} else {
		resp := &collectionpb.TokenResponse{
			TokenRes: "NO",
		}
		return resp, nil
	}
}

//Check Token
//-------------------------
//send data
func (*server) SendData(stream collectionpb.CollectionService_SendDataServer) error {
	log.Println("Send data...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &collectionpb.DataStreamResponse{
				DataRes: "Done",
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatal("Err while Recv Average %v", err)
			return err
		}
		log.Println("Receive num %v", req)
		//send to kafka
		//fmt.Printf("%T", req.GetDataReq())
		service.ProducerMessage(req.String(), "testTopic")
		//send to kafka
	}

}

//send data
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")
	if err != nil {
		log.Fatal("Error while create listen %v", err)
	}
	s := grpc.NewServer()

	collectionpb.RegisterCollectionServiceServer(s, &server{}) // ham nay tu protoc go tao ra

	fmt.Println("Collection data is running...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Error while server %v", err)
	}
}

//check token
func CheckToken(token string) string {
	tokenInputString := `{"token":"` + token + `"}`
	//fmt.Println(tokenInputString)
	//-------------------parse string to Token
	////////var tokenInput TokenInput
	////////json.Unmarshal([]byte(tokenInputString), &tokenInput)
	////////fmt.Println(tokenInput.Token)
	//-------------------parse string to Token
	url := "http://192.168.0.106:1323/api/sso/parseToken"
	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "")
	s.Header = h
	var jsonStr = []byte(tokenInputString)

	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(resp.RawText()), &result)
	if result["code"].(string) == "0" {
		tokenInfo := result["data"].(map[string]interface{})
		status := (tokenInfo["userstatus"]).(string)
		return status
	} else {
		return ""
	}
	//fmt.Println(result["data"])
	//
	/*tokenInfo := result["data"].(map[string]interface{})*/
	/*for key, value := range tokenInfo {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}*/
	/*status := (tokenInfo["userstatus"]).(string)
	return status*/
}
