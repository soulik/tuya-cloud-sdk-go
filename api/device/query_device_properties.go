package device

import (
	"fmt"

	"github.com/soulik/tuya-cloud-sdk-go/api/common"
)

type QueryDevicePropertiesReq struct {
	DeviceID string
	Codes    string `json:"codes,omitempty"`
}

func (t *QueryDevicePropertiesReq) Method() string {
	return common.RequestGet
}

func (t *QueryDevicePropertiesReq) API() string {
	if t.Codes != "" {
		return fmt.Sprintf("/v2.0/cloud/thing/%s/shadow/properties?codes=%s", t.DeviceID, t.Codes)
	} else {
		return fmt.Sprintf("/v2.0/cloud/thing/%s/shadow/properties", t.DeviceID)
	}
}

// QueryDeviceProperties Based on the device ID, query the property status reported by the device to the cloud.
func QueryDeviceProperties(deviceID string, args ...string) (*QueryDevicePropertiesResponse, error) {
	a := &QueryDevicePropertiesReq{DeviceID: deviceID}
	if len(args) > 0 {
		a.Codes = args[0]
	}
	resp := &QueryDevicePropertiesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type Property struct {
	Code       string      `json:"code"`
	Value      interface{} `json:"value"`
	DpId       int         `json:"dp_id"`
	Time       int         `json:"time"`
	Type       string      `json:"type"`
	CustomName string      `json:"custom_name"`
}

type QueryDevicePropertiesResponse struct {
	Success bool       `json:"success"`
	T       int64      `json:"t"`
	Result  []Property `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
