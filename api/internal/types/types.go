// Code generated by goctl. DO NOT EDIT.
package types

type OSSCallBackReq struct {
	Bucket    string `json:"bucket,optional"`
	Object    string `json:"object,optional"`
	ClientIP  string `json:"clientIp,optional"`
	ReqId     string `json:"reqId,optional"`
	Operation string `json:"operation,optional"`
}

type OSSCallBackResp struct {
	Bucket    string `json:"bucket"`
	Object    string `json:"object"`
	ClientIP  string `json:"clientIp"`
	ReqId     string `json:"reqId"`
	Operation string `json:"operation"`
}
