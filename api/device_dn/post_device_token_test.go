package device_dn

import (
	"github.com/soulik/tuya-cloud-sdk-go/api/common"

	"testing"
)

func TestGetDeviceToken(t *testing.T) {
	d := &PostDeviceTokenReq{
		UID:        common.Ed.TestDataUID,
		TimeZoneID: common.Ed.TestDataTimeZoneID,
	}
	resp, err := PostDeviceToken(d)
	if err != nil {
		t.Error("err:", err)
		return
	}
	if !resp.Success {
		t.Error("err:", resp.Msg)
	}
	t.Logf("%+v\n", resp)
}
