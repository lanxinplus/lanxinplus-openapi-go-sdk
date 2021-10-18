package sdk

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

//注：以下配置信息从蓝信 开发者中心-应用开发-基本信息 中获取
const (
	SERVER     = "" // 蓝信开放平台网关地址, e.g.: https://example.com/open/apigw
	APP_ID     = "" // 应用ID, e.g.: 1234567-7654321
	APP_SECRET = "" // 应用密钥, e.g.: D25F65E65D887AEFD9C92B00310286FA
)

var (
	client *ClientWithResponses
)

func init() {
	c, err := NewClientWithResponses(SERVER)
	if err != nil {
		panic(err)
	}
	client = c
}

func GetV1AppToken() string {
	params := V1AppTokenCreateParams{}
	params.SetGrantType("client_credential").
		SetAppid(APP_ID).
		SetSecret(APP_SECRET)

	resp, err := client.V1AppTokenCreateWithResponse(context.TODO(), &params)
	if err != nil {
		panic(err)
	}
	if resp == nil ||
		resp.GetErrCode() != 0 ||
		resp.GetData() == nil ||
		resp.GetData().GetAppToken() == "" {
		panic("invalid app_token")
	}
	return resp.GetData().GetAppToken()
}

// /v1/app/roaming/orgs/fetch
func Test_V1AppRoamingOrgsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V1AppRoamingOrgsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1AppRoamingOrgsFetchParams{}
	params.SetAppToken(GetV1AppToken()).
		SetPageSize(100).
		SetBaseVersion("")

	tests := []struct {
		name string
		args args

		want *V1AppRoamingOrgsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V1AppRoamingOrgsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1AppRoamingOrgsFetchWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1AppRoamingOrgsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1AppRoamingOrgsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/apptoken/create
func Test_V1AppTokenCreateWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V1AppTokenCreateParams

		reqEditors []RequestEditorFn
	}
	params := V1AppTokenCreateParams{}
	params.SetGrantType("").
		SetAppid("").
		SetSecret("")

	tests := []struct {
		name string
		args args

		want *V1AppTokenCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V1AppTokenCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1AppTokenCreateWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1AppTokenCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1AppTokenCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/chat/notification/{userid}/fetch
func Test_V1ChatNotificationFetchWithResponse(t *testing.T) {
	type args struct {
		ctx    context.Context
		userid string
		params *V1ChatNotificationFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1ChatNotificationFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1ChatNotificationFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				userid: "string",
				params: &params,
			},

			want: &V1ChatNotificationFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1ChatNotificationFetchWithResponse(tt.args.ctx, tt.args.userid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1ChatNotificationFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1ChatNotificationFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/create
func Test_V1DeptsCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1DeptsCreateParams
		body       V1DeptsCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1DeptsCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1DeptsCreateRequestBody{}
	body.SetExternalId("").
		SetName("").
		SetOrderNumber(0).
		SetParentId("").
		SetTags([]*string{})

	tests := []struct {
		name string
		args args

		want *V1DeptsCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1DeptsCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/children/fetch
func Test_V1DeptsChildrenFetchWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		params       *V1DeptsChildrenFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsChildrenFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsChildrenFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				params:       &params,
			},

			want: &V1DeptsChildrenFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsChildrenFetchWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsChildrenFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsChildrenFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/delete
func Test_V1DeptsDeleteWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		params       *V1DeptsDeleteParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				params:       &params,
			},

			want: &V1DeptsDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsDeleteWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/fetch
func Test_V1DeptsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		params       *V1DeptsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				params:       &params,
			},

			want: &V1DeptsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsFetchWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/staffs/fetch
func Test_V1DeptsStaffsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		params       *V1DeptsStaffsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsStaffsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsStaffsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				params:       &params,
			},

			want: &V1DeptsStaffsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsStaffsFetchWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsStaffsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsStaffsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/staffs/{staffid}/create
func Test_V1DeptsStaffsCreateWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		staffid      string
		params       *V1DeptsStaffsCreateParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsStaffsCreateParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsStaffsCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				staffid:      "string",
				params:       &params,
			},

			want: &V1DeptsStaffsCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsStaffsCreateWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsStaffsCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsStaffsCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/staffs/{staffid}/delete
func Test_V1DeptsStaffsDeleteWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		staffid      string
		params       *V1DeptsStaffsDeleteParams

		reqEditors []RequestEditorFn
	}
	params := V1DeptsStaffsDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1DeptsStaffsDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				staffid:      "string",
				params:       &params,
			},

			want: &V1DeptsStaffsDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsStaffsDeleteWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsStaffsDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsStaffsDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/departments/{departmentid}/update
func Test_V1DeptsUpdateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		departmentid string
		params       *V1DeptsUpdateParams
		body         V1DeptsUpdateRequestBody
		reqEditors   []RequestEditorFn
	}
	params := V1DeptsUpdateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1DeptsUpdateRequestBody{}
	body.SetName("").
		SetOrderNumber(0).
		SetParentId("").
		SetTags([]*string{})

	tests := []struct {
		name string
		args args

		want *V1DeptsUpdateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				departmentid: "string",
				params:       &params,
				body:         body,
			},

			want: &V1DeptsUpdateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1DeptsUpdateWithBodyWithResponse(tt.args.ctx, tt.args.departmentid, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1DeptsUpdateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1DeptsUpdateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/jsapitoken/create
func Test_V1JsApiTokenCreateWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V1JsApiTokenCreateParams

		reqEditors []RequestEditorFn
	}
	params := V1JsApiTokenCreateParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1JsApiTokenCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V1JsApiTokenCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1JsApiTokenCreateWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1JsApiTokenCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1JsApiTokenCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/medias/create
func Test_V1MediasCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params      *V1MediasCreateParams
		contentType string
		body        io.Reader
		reqEditors  []RequestEditorFn
	}
	params := V1MediasCreateParams{}
	params.SetAppToken(GetV1AppToken()).
		SetType("")

	tests := []struct {
		name string
		args args

		want *V1MediasCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params:      &params,
				contentType: "", body: nil,
			},

			want: &V1MediasCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MediasCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.contentType, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MediasCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MediasCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/medias/{mediaid}/fetch
func Test_V1MediasFetchWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		mediaid string
		params  *V1MediasFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1MediasFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *http.Response

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				mediaid: "string",
				params:  &params,
			},

			want: &http.Response{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MediasFetchWithResponse(tt.args.ctx, tt.args.mediaid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MediasFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MediasFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {

				body, err := ioutil.ReadAll(got.Body)
				if err != nil {
					t.Errorf("ClientWithResponses.V1MediasFetchWithResponse() = %v, read body err %v", got, err)
				}
				filename := strings.Replace(got.Header.Get("Content-disposition"), "attachment;filename=", "", -1)
				tmpFile, err := os.Create("./" + filename)
				if err != nil {
					t.Errorf("ClientWithResponses.V1MediasFetchWithResponse() = %v, create tmp file err %v", got, err)
				}
				defer tmpFile.Close()
				tmpFile.Write(body)

			}
		})
	}
}

// /v1/medias/{mediaid}/path/fetch
func Test_V1MediasPathFetchWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		mediaid string
		params  *V1MediasPathFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1MediasPathFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1MediasPathFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				mediaid: "string",
				params:  &params,
			},

			want: &V1MediasPathFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MediasPathFetchWithResponse(tt.args.ctx, tt.args.mediaid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MediasPathFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MediasPathFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/messages/create
func Test_V1MessagesCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1MessagesCreateParams
		body       V1MessagesCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1MessagesCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1MessagesCreateRequestBody{}
	body.SetAccountId("").
		SetAttach("").
		SetDepartmentIdList([]*string{}).
		SetEntryId("").
		SetMsgData(nil).
		SetMsgType("").
		SetUserIdList([]*string{})

	tests := []struct {
		name string
		args args

		want *V1MessagesCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1MessagesCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MessagesCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MessagesCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MessagesCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/messages/notification/create
func Test_V1MessagesNotificationCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1MessagesNotificationCreateParams
		body       V1MessagesNotificationCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1MessagesNotificationCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1MessagesNotificationCreateRequestBody{}
	body.SetMsgData(nil).
		SetMsgType("").
		SetReceiverIds([]*string{}).
		SetSenderId("").
		SetUuid("")

	tests := []struct {
		name string
		args args

		want *V1MessagesNotificationCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1MessagesNotificationCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MessagesNotificationCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MessagesNotificationCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MessagesNotificationCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/messages/revoke
func Test_V1MessagesRevokeWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1MessagesRevokeParams
		body       V1MessagesRevokeRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1MessagesRevokeParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1MessagesRevokeRequestBody{}
	body.SetChatType("").
		SetMessageIds([]*string{}).
		SetSenderId("").
		SetSysMsg(SystemMsg{})

	tests := []struct {
		name string
		args args

		want *V1MessagesRevokeResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1MessagesRevokeResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1MessagesRevokeWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1MessagesRevokeWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1MessagesRevokeWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/org/{orgid}/extrafieldids/fetch
func Test_V1OrgExtraFieldIdsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx    context.Context
		orgid  string
		params *V1OrgExtraFieldIdsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1OrgExtraFieldIdsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1OrgExtraFieldIdsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				orgid:  "string",
				params: &params,
			},

			want: &V1OrgExtraFieldIdsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1OrgExtraFieldIdsFetchWithResponse(tt.args.ctx, tt.args.orgid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1OrgExtraFieldIdsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1OrgExtraFieldIdsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/org/{orgid}/fetch
func Test_V1OrgFetchWithResponse(t *testing.T) {
	type args struct {
		ctx    context.Context
		orgid  string
		params *V1OrgFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1OrgFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1OrgFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				orgid:  "string",
				params: &params,
			},

			want: &V1OrgFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1OrgFetchWithResponse(tt.args.ctx, tt.args.orgid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1OrgFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1OrgFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/org/{orgid}/role/create
func Test_V1OrgRoleCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		orgid      string
		params     *V1OrgRoleCreateParams
		body       V1OrgRoleCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1OrgRoleCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1OrgRoleCreateRequestBody{}
	body.SetCreator("").
		SetDescription("").
		SetRoleName("")

	tests := []struct {
		name string
		args args

		want *V1OrgRoleCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				orgid:  "string",
				params: &params,
				body:   body,
			},

			want: &V1OrgRoleCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1OrgRoleCreateWithBodyWithResponse(tt.args.ctx, tt.args.orgid, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1OrgRoleCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1OrgRoleCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/org/{orgid}/role/{roleid}/members/fetch
func Test_V1OrgRoleMembersFetchWithResponse(t *testing.T) {
	type args struct {
		ctx    context.Context
		orgid  string
		roleid string
		params *V1OrgRoleMembersFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1OrgRoleMembersFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1OrgRoleMembersFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				orgid:  "string",
				roleid: "string",
				params: &params,
			},

			want: &V1OrgRoleMembersFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1OrgRoleMembersFetchWithResponse(tt.args.ctx, tt.args.orgid, tt.args.roleid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1OrgRoleMembersFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1OrgRoleMembersFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/rolemember/create
func Test_V1RoleMemberCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1RoleMemberCreateParams
		body       V1RoleMemberCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1RoleMemberCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1RoleMemberCreateRequestBody{}
	body.SetCreator("").
		SetRoleId("").
		SetStaffIds([]*string{})

	tests := []struct {
		name string
		args args

		want *V1RoleMemberCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1RoleMemberCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1RoleMemberCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1RoleMemberCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1RoleMemberCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/rolemember/{rolememberid}/delete
func Test_V1RoleMemberDeleteWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx          context.Context
		rolememberid string
		params       *V1RoleMemberDeleteParams
		body         V1RoleMemberDeleteRequestBody
		reqEditors   []RequestEditorFn
	}
	params := V1RoleMemberDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1RoleMemberDeleteRequestBody{}
	body.SetOperator("")

	tests := []struct {
		name string
		args args

		want *V1RoleMemberDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:          context.TODO(),
				rolememberid: "string",
				params:       &params,
				body:         body,
			},

			want: &V1RoleMemberDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1RoleMemberDeleteWithBodyWithResponse(tt.args.ctx, tt.args.rolememberid, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1RoleMemberDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1RoleMemberDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/create
func Test_V1StaffsCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1StaffsCreateParams
		body       V1StaffsCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1StaffsCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1StaffsCreateRequestBody{}
	body.SetAddress("").
		SetAvatarId("").
		SetAvatarUrl("").
		SetBirthdate("").
		SetCareer([]*Resume{}).
		SetDepartments([]*UDepartment{}).
		SetDisablePasswordReset(false).
		SetDuties("").
		SetEducation([]*Resume{}).
		SetEmail("").
		SetEmployeeNumber("").
		SetExtAttr1("").
		SetExtAttr2("").
		SetExternalId("").
		SetExtraFieldSet(map[string]string{}).
		SetExtraPhones([]*MobilePhone{}).
		SetGender(0).
		SetIdNumber("").
		SetIntroduction(Introduction{}).
		SetLoginName("").
		SetLoginPassword("").
		SetLoginWays([]*int{}).
		SetMobilePhone(MobilePhone{}).
		SetName("").
		SetNationality("").
		SetNativePlace("").
		SetOrgId("").
		SetParties("").
		SetSignature("").
		SetSmsInvitation(false).
		SetStatus(0).
		SetTags([]*string{})

	tests := []struct {
		name string
		args args

		want *V1StaffsCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1StaffsCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/{staffid}/delete
func Test_V1StaffsDeleteWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		staffid string
		params  *V1StaffsDeleteParams

		reqEditors []RequestEditorFn
	}
	params := V1StaffsDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1StaffsDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				staffid: "string",
				params:  &params,
			},

			want: &V1StaffsDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsDeleteWithResponse(tt.args.ctx, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/{staffid}/departmentancestors/fetch
func Test_V1StaffsDeptAncestorsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		staffid string
		params  *V1StaffsDeptAncestorsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1StaffsDeptAncestorsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1StaffsDeptAncestorsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				staffid: "string",
				params:  &params,
			},

			want: &V1StaffsDeptAncestorsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsDeptAncestorsFetchWithResponse(tt.args.ctx, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsDeptAncestorsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsDeptAncestorsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/{staffid}/fetch
func Test_V1StaffsFetchWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		staffid string
		params  *V1StaffsFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1StaffsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1StaffsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				staffid: "string",
				params:  &params,
			},

			want: &V1StaffsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsFetchWithResponse(tt.args.ctx, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/{staffid}/infor/fetch
func Test_V1StaffsInforFetchWithResponse(t *testing.T) {
	type args struct {
		ctx     context.Context
		staffid string
		params  *V1StaffsInforFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1StaffsInforFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1StaffsInforFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				staffid: "string",
				params:  &params,
			},

			want: &V1StaffsInforFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsInforFetchWithResponse(tt.args.ctx, tt.args.staffid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsInforFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsInforFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/staffs/{staffid}/update
func Test_V1StaffsUpdateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		staffid    string
		params     *V1StaffsUpdateParams
		body       V1StaffsUpdateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1StaffsUpdateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1StaffsUpdateRequestBody{}
	body.SetAddress("").
		SetAvatarId("").
		SetBirthdate("").
		SetCareer([]*Resume{}).
		SetDepartments([]*UDepartment{}).
		SetDuties("").
		SetEducation([]*Resume{}).
		SetEmail("").
		SetEmployeeNumber("").
		SetExtAttr1("").
		SetExtAttr2("").
		SetExternalId("").
		SetExtraFieldSet(map[string]interface{}{}).
		SetExtraPhones([]*MobilePhone{}).
		SetGender(0).
		SetIdNumber("").
		SetIntroduction(Introduction{}).
		SetLoginName("").
		SetLoginWays([]*int{}).
		SetMobilePhone(MobilePhone{}).
		SetName("").
		SetNationality("").
		SetNativePlace("").
		SetParties("").
		SetSignature("").
		SetStatus(0).
		SetTags([]*string{})

	tests := []struct {
		name string
		args args

		want *V1StaffsUpdateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:     context.TODO(),
				staffid: "string",
				params:  &params,
				body:    body,
			},

			want: &V1StaffsUpdateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1StaffsUpdateWithBodyWithResponse(tt.args.ctx, tt.args.staffid, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1StaffsUpdateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1StaffsUpdateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/taggroups/create
func Test_V1TagGroupsCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1TagGroupsCreateParams
		body       V1TagGroupsCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagGroupsCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagGroupsCreateRequestBody{}
	body.SetOrgId("").
		SetTagGroupCategory(0).
		SetTagGroupName("").
		SetTagGroupType("")

	tests := []struct {
		name string
		args args

		want *V1TagGroupsCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1TagGroupsCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagGroupsCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagGroupsCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagGroupsCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/taggroups/fetch
func Test_V1TagGroupsFetchWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1TagGroupsFetchParams
		body       V1TagGroupsFetchRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagGroupsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagGroupsFetchRequestBody{}
	body.SetCategoryList([]*int{}).
		SetOrgId("").
		SetTagGroupType("")

	tests := []struct {
		name string
		args args

		want *V1TagGroupsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1TagGroupsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagGroupsFetchWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagGroupsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagGroupsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/taggroups/{tag_group_id}/delete
func Test_V1TagGroupsDeleteWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		tagGroupId string
		params     *V1TagGroupsDeleteParams

		reqEditors []RequestEditorFn
	}
	params := V1TagGroupsDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1TagGroupsDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        context.TODO(),
				tagGroupId: "string",
				params:     &params,
			},

			want: &V1TagGroupsDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagGroupsDeleteWithResponse(tt.args.ctx, tt.args.tagGroupId, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagGroupsDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagGroupsDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/taggroups/{tag_group_id}/fetch
func Test_V1TagGroupsInfoFetchWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		tagGroupId string
		params     *V1TagGroupsInfoFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1TagGroupsInfoFetchParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1TagGroupsInfoFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        context.TODO(),
				tagGroupId: "string",
				params:     &params,
			},

			want: &V1TagGroupsInfoFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagGroupsInfoFetchWithResponse(tt.args.ctx, tt.args.tagGroupId, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagGroupsInfoFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagGroupsInfoFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/taggroups/{tag_group_id}/update
func Test_V1TagGroupsUpdateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		tagGroupId string
		params     *V1TagGroupsUpdateParams
		body       V1TagGroupsUpdateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagGroupsUpdateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagGroupsUpdateRequestBody{}
	body.SetTagGroupName("")

	tests := []struct {
		name string
		args args

		want *V1TagGroupsUpdateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:        context.TODO(),
				tagGroupId: "string",
				params:     &params,
				body:       body,
			},

			want: &V1TagGroupsUpdateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagGroupsUpdateWithBodyWithResponse(tt.args.ctx, tt.args.tagGroupId, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagGroupsUpdateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagGroupsUpdateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/tags/create
func Test_V1TagsCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1TagsCreateParams
		body       V1TagsCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagsCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagsCreateRequestBody{}
	body.SetTagGroupId("").
		SetTagName("")

	tests := []struct {
		name string
		args args

		want *V1TagsCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1TagsCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagsCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagsCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagsCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/tags/meta/fetch
func Test_V1TagsMetaFetchWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1TagsMetaFetchParams
		body       V1TagsMetaFetchRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagsMetaFetchParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagsMetaFetchRequestBody{}
	body.SetTagIds([]*string{})

	tests := []struct {
		name string
		args args

		want *V1TagsMetaFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1TagsMetaFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagsMetaFetchWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagsMetaFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagsMetaFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/tags/staffids/fetch
func Test_V1TagsFetchWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V1TagsFetchParams
		body       V1TagsFetchRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagsFetchParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagsFetchRequestBody{}
	body.SetOrgId("").
		SetTagFilters([]*TagFilter{})

	tests := []struct {
		name string
		args args

		want *V1TagsFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V1TagsFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagsFetchWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagsFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagsFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/tags/{tagid}/delete
func Test_V1TagsDeleteWithResponse(t *testing.T) {
	type args struct {
		ctx    context.Context
		tagid  string
		params *V1TagsDeleteParams

		reqEditors []RequestEditorFn
	}
	params := V1TagsDeleteParams{}
	params.SetAppToken(GetV1AppToken())

	tests := []struct {
		name string
		args args

		want *V1TagsDeleteResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				tagid:  "string",
				params: &params,
			},

			want: &V1TagsDeleteResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagsDeleteWithResponse(tt.args.ctx, tt.args.tagid, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagsDeleteWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagsDeleteWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/tags/{tagid}/update
func Test_V1TagsUpdateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx        context.Context
		tagid      string
		params     *V1TagsUpdateParams
		body       V1TagsUpdateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V1TagsUpdateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V1TagsUpdateRequestBody{}
	body.SetTagName("")

	tests := []struct {
		name string
		args args

		want *V1TagsUpdateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx:    context.TODO(),
				tagid:  "string",
				params: &params,
				body:   body,
			},

			want: &V1TagsUpdateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1TagsUpdateWithBodyWithResponse(tt.args.ctx, tt.args.tagid, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1TagsUpdateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1TagsUpdateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/users/fetch
func Test_V1UsersFetchWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V1UsersFetchParams

		reqEditors []RequestEditorFn
	}
	params := V1UsersFetchParams{}
	params.SetAppToken(GetV1AppToken()).
		SetUserToken("")

	tests := []struct {
		name string
		args args

		want *V1UsersFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V1UsersFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1UsersFetchWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1UsersFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1UsersFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v1/usertoken/create
func Test_V1UserTokenCreateWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V1UserTokenCreateParams

		reqEditors []RequestEditorFn
	}
	params := V1UserTokenCreateParams{}
	params.SetAppToken(GetV1AppToken()).
		SetGrantType("").
		SetCode("")

	tests := []struct {
		name string
		args args

		want *V1UserTokenCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V1UserTokenCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V1UserTokenCreateWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V1UserTokenCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V1UserTokenCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v2/chat/notification/update
func Test_V2ChatNotificationUpdateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V2ChatNotificationUpdateParams
		body       V2ChatNotificationUpdateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V2ChatNotificationUpdateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V2ChatNotificationUpdateRequestBody{}
	body.SetStatusList([]*StatusInfo{})

	tests := []struct {
		name string
		args args

		want *V2ChatNotificationUpdateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V2ChatNotificationUpdateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V2ChatNotificationUpdateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V2ChatNotificationUpdateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V2ChatNotificationUpdateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v2/event/notification/create
func Test_V2EventNotificationCreateWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V2EventNotificationCreateParams
		body       V2EventNotificationCreateRequestBody
		reqEditors []RequestEditorFn
	}
	params := V2EventNotificationCreateParams{}
	params.SetAppToken(GetV1AppToken())

	body := V2EventNotificationCreateRequestBody{}
	body.SetEvents([]*Event{})

	tests := []struct {
		name string
		args args

		want *V2EventNotificationCreateResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V2EventNotificationCreateResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V2EventNotificationCreateWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V2EventNotificationCreateWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V2EventNotificationCreateWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v2/staffs/id_mapping/fetch
func Test_V2StaffsIdMappingFetchWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params *V2StaffsIdMappingFetchParams

		reqEditors []RequestEditorFn
	}
	params := V2StaffsIdMappingFetchParams{}
	params.SetAppToken(GetV1AppToken()).
		SetOrgId("").
		SetIdType("").
		SetIdValue("")

	tests := []struct {
		name string
		args args

		want *V2StaffsIdMappingFetchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
			},

			want: &V2StaffsIdMappingFetchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V2StaffsIdMappingFetchWithResponse(tt.args.ctx, tt.args.params, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V2StaffsIdMappingFetchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V2StaffsIdMappingFetchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}

// /v2/staffs/search
func Test_V2StaffsSearchWithBodyWithResponse(t *testing.T) {
	type args struct {
		ctx context.Context

		params     *V2StaffsSearchParams
		body       V2StaffsSearchRequestBody
		reqEditors []RequestEditorFn
	}
	params := V2StaffsSearchParams{}
	params.SetAppToken(GetV1AppToken()).
		SetUserId("")

	body := V2StaffsSearchRequestBody{}
	body.SetKeyword("").
		SetRecursive(false).
		SetSearchScope(SearchScope{})

	tests := []struct {
		name string
		args args

		want *V2StaffsSearchResponse

		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.TODO(),

				params: &params,
				body:   body,
			},

			want: &V2StaffsSearchResponse{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := client.V2StaffsSearchWithBodyWithResponse(tt.args.ctx, tt.args.params, tt.args.body, tt.args.reqEditors...)

			if (err != nil) != tt.wantErr {
				t.Errorf("ClientWithResponses.V2StaffsSearchWithResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientWithResponses.V2StaffsSearchWithResponse() = %v, want %v", got, tt.want)
			}
			if got != nil {
				fmt.Println(got.ToString())
			}
		})
	}
}
