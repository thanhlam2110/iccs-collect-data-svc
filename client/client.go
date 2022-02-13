package main

import (
	"context"
	"fmt"
	"log"

	collectionpb "github.com/thanhlam/home-collect-data-svc/collectionpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50070", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Err while dial $v", err)
	}
	defer cc.Close()
	client := collectionpb.NewCollectionServiceClient(cc)
	tokenResult := callCheckToken(client)
	if tokenResult == "YES" {
		//fmt.Println("YEAHHHH")
		callSendData(client)
	}
}

//check token
func callCheckToken(c collectionpb.CollectionServiceClient) string {
	log.Println("Call Check Token")
	resp, err := c.CheckToken(context.Background(), &collectionpb.TokenRequest{
		TokenReq: "eyJ6aXAiOiJERUYiLCJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiY3R5IjoiSldUIiwidHlwIjoiSldUIiwib3JnLmFwZXJlby5jYXMuc2VydmljZXMuUmVnaXN0ZXJlZFNlcnZpY2UiOiIzIn0..QtJDoXtlwcWPrVTU85wy8w.zKszaWnNUc02-Nrp_H85QW3pYSmXEZi0djej9OZU_52pwWIxJqW_kTVPjSIrQ6Cf-Cun0lEhJAHFj7iVwY3yxgNg_tEJLtZRuNk19pUSyMVMAAdQJSDwcSRsQlvTY6BmtyLlOB-KC8a2LoMv28YH-8HszpePRRjjqxiFCqpclEBy8OHTqpUwQz4i584zC11UZAs4BEu0eSk_cpnu0i1d0mkrUeBztnrxkcrECUaej-5m1zfbcI3ErErxyUBDjBNxHN0DofYdw5ZhJRxnHAhBRl55V3f-S3r2Uii6mATbEXeLCx9VtQwUM0IxLgHEYhBq4HpXoESjn-IFebxQlaFJ9Wq4jUNFGmnLlsWzfSSLE8muC9iZ_miVozCWnnIMcKwOeGNQytB01Nuj2qVfR6l7DwLj86dDLt5b_4UXVBLvdxV-qazDB5E9jkXGhnTUans6TdanrYnrgFvU7cj48YBp7wENJoU9yaetJsv5Gi70CgcahKoxhYmGX-QYf2oZmpxYRR6B22v0N-kQz9uB-QaJCC-7whYn43JjSRJxUyDqKNJYAGZSUfvOhSPeg5ea70KXLkOe7vEiz2skIkRFm2fbUwqLXkncFz__geERwtfHHld9U6_8baEfX4Dq7hzadz3jBCIG_rE6Mjri6AhU5smxbDL5174I1UaQGWT-hhVl9LP0rSFrIKOUJ_ABLE9O2NDBxBQI4kmP8TI5juhTDqHiDcoL8bHrWb94w-xLZBC2Ju8Pyw80XTwInNa0W3RgB4y4.c_kT0NCIi0uYLncw7HTQPQ",
		//TokenReq: "ThanhLamTokenFake",
	})
	if err != nil {
		log.Fatal("Call check token wrong %v", err)
	}
	log.Println("Call check token response", resp.GetTokenRes())
	return resp.GetTokenRes()
}

//send data
func callSendData(c collectionpb.CollectionServiceClient) {
	log.Println("Call send data")
	stream, err := c.SendData(context.Background())
	if err != nil {
		log.Fatal("Call send data err %v", err)
	}
	//listReq := []collectionpb.DataStreamRequest{}
	var listReq []collectionpb.DataStreamRequest
	var i int
	for i = 1; i <= 10; i++ {
		j := collectionpb.DataStreamRequest{
			DataReq: int32(i),
		}
		listReq = append(listReq, j)
	}
	//gửi đi
	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatal("Send data request err %v", err)
		}
	}
	//gửi hết thì nhận response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Receive send response err %v", err)
	}
	fmt.Printf("Response %v", resp)

}
