package internal

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/imagesearch"
)

func main() {
	endpoints.AddEndpointMapping("<region>", "ImageSearch", "imagesearch.<region>.aliyuncs.com")
	// 创建client实例
	client, err := imagesearch.NewClientWithAccessKey(
		"<region>",                 // 您的可用区ID
		"<your-access-key-id>",     // 您的Access Key ID
		"<your-access-key-secret>") // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}
	// 查询图片
	searchRequest := imagesearch.CreateSearchImageRequest()
	searchRequest.InstanceName = "demo"
	searchRequest.PicName = "test"
	searchRequest.ProductId = "test"
	searchRequest.Type = "SearchByName"
	searchResponse, err := client.SearchImage(searchRequest)
	if err != nil {
		// 异常处理
		panic(err)
	}
	fmt.Println(searchResponse)
}
