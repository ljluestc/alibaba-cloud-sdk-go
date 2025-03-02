package gtm

import (
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeGtmInstanceRequest struct {
    requests.RpcRequest
    InstanceId string `position:"Query" name:"InstanceId"`
}

type DescribeGtmInstanceResponse struct {
    responses.BaseResponse
    RequestId string `json:"RequestId" xml:"RequestId"`
    Instance  struct {
        InstanceId   string `json:"InstanceId" xml:"InstanceId"`
        InstanceName string `json:"InstanceName" xml:"InstanceName"`
    } `json:"Instance" xml:"Instance"`
}

func (client *Client) DescribeGtmInstance(request *DescribeGtmInstanceRequest) (response *DescribeGtmInstanceResponse, err error) {
    response = &DescribeGtmInstanceResponse{}
    err = client.DoAction(request, response)
    return
}

type UpdateGtmInstanceRequest struct {
    requests.RpcRequest
    InstanceId   string `position:"Query" name:"InstanceId"`
    InstanceName string `position:"Query" name:"InstanceName"`
}

type UpdateGtmInstanceResponse struct {
    responses.BaseResponse
    RequestId string `json:"RequestId" xml:"RequestId"`
}

func (client *Client) UpdateGtmInstance(request *UpdateGtmInstanceRequest) (response *UpdateGtmInstanceResponse, err error) {
    response = &UpdateGtmInstanceResponse{}
    err = client.DoAction(request, response)
    return
}