package client

import (
	"fmt"
	"log"
	"net/rpc"

	contract "chapter1/rpc/contract"
)

const port = 1234

func CreateClient() *rpc.Client {
	// Dial: hội thoại.
	// Questions: Khi đọc đến đây, ta sẽ thấy một câu hỏi lớn đó là, làm sao mà CreateClient() nó hoạt động, mà không có chỗ nào gọi nó?
	// Answer: Để trả lời câu hỏi trên? ta phải đọc phần comment dưới và tìm hiểu thêm: https://pkg.go.dev/encoding/gob@go1.17.6#pkg-index
	/*
		With gobs, the source and destination values and types do not need to correspond exactly, when you send struct, if a field is in the source but not
		in the receiving struct, then the decoder will ignore this field and the processing will continue without error. If a field is present in the destination
		that is not in the source, then again the decoder will ignore this field and will successfully process the rest of the message.
	*/
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	// After connect to server successfully, tức có nghĩa, khi mà ta sử dụng client.Call(), thì nó đã auto chạy được Dial() hay gì rồi?
	// Sau đó, ta có thể make a request đến Server bằng việc input các parameter vào client.Call().
	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
