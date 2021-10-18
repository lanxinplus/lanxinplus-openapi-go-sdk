# 蓝信 Go SDK

## 关于
> - 此Go SDK基于[蓝信开放平台API]构建
> - 蓝信开放平台，便于企业应用与蓝信集成，让协同与管理更加高效
> - 蓝信开放接口SDK，便捷调用服务端API，例如：认证授权、通讯录、消息通知、角色权限、媒体素材、组织管理、应用管理等具体可以访问 [蓝信开放平台API] 文档 看看

## 运行环境
> - Go 1.13及以上

## 安装方法
#### 通过 GitHub 安装 SDK
> - 运行`go get github.com/lanxinplus/lanxinplus-openapi-go-sdk/sdk` 命令获取远程代码包
> - 在您的代码中使用 `import "github.com/lanxinplus/lanxinplus-openapi-go-sdk/sdk" `来引入 Go SDK包

## 快速使用
#### 获取 APP_TOKEN
```
import (
	"context"
	"fmt"
	
	"github.com/lanxinplus/lanxinplus-openapi-go-sdk/sdk"
)

//注：以下配置信息从蓝信 开发者中心-应用开发-基本信息 中获取
const (
	SERVER     = "" // 蓝信开放平台网关地址, e.g.: https://example.com/open/apigw
	APP_ID     = "" // 应用ID, e.g.: 1234567-7654321
	APP_SECRET = "" // 应用密钥, e.g.: D25F65E65D887AEFD9C92B00310286FA
)

func Example() {
	client, err := sdk.NewClientWithResponses(SERVER)
	if err != nil {
		// HandleError(err)
	}

	params := sdk.V1AppTokenCreateParams{}
	params.SetGrantType("client_credential").
		SetAppid(APP_ID).
		SetSecret(APP_SECRET)

	// 返回内部处理的 Response
	resp, err := client.V1AppTokenCreateWithResponse(context.TODO(), &params)
	if err != nil {
		// HandleError(err)
	}
	fmt.Println(resp.ToString())
	fmt.Println(resp.GetErrCode())
	fmt.Println(resp.GetData().GetAppToken())
	fmt.Println(resp.GetData().GetExpiresIn())
}
```

## 测试用例使用说明
### 运行test
> - 拷贝示例文件到Go SDK的安装路径（即GOPATH变量中的第一个路径），进入Go SDK的代码目录`src/github.com/lanxinplus/lanxinplus-openapi-go-sdk/sdk`
> - 修改sdk_test.go文件const里的SERVER、APP_ID、APP_SECRET 配置
> - 请在您的工程目录下执行`go test -run Test_XXX`

## 联系我们
- [蓝信官方网站](https://www.lanxin.cn/)
- [蓝信开放平台文档中心](https://openapi.lanxin.cn/doc/#/)

[蓝信开放平台API]: https://openapi.lanxin.cn/doc/#/server-api/
