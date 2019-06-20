# wechat 微信统一服务
### 微服务内核访问演示
```
{
    "service": "wechat",
    "method": "Wechat.ProcessCommonRequest",
    "request": {
        "Domain": "miniprogram",
        "ApiName": "code2Session",
        "QueryParams": {
        	"PhoneNumbers":"13954386521",
        	"SignName":"微信",
        	"TemplateCode":"SMS_135275049",
        	"TemplateParam":"{code: '562374'}"
        }
    }
}
```
- 具体参数参考微信开发文档
- https://api.wechat.com/