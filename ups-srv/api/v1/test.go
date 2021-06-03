package v1

import (
	"ups-srv/pkg/invoker"
	"waf-srv/model"

	"github.com/gotomicro/ego/core/elog"

	"github.com/iGoogle-ink/gopay"

	"github.com/iGoogle-ink/gopay/alipay"

	"github.com/hehanpeng/gofund/common/global/api"

	"github.com/gin-gonic/gin"
)

// 不能依赖waf 只是为了省事
func Hello(c *gin.Context) {
	var ttoInfo model.TtoInfo
	_ = c.ShouldBindJSON(&ttoInfo)
	api.Ok(c)
}

func TestAli(c *gin.Context) {
	//  初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	privateKey := "MIIEowIBAAKCAQEAhHgqVG0SSEVgMGqlQH+8McIbDtrCJlLNMmUjAJYrIvC4hoz+euWuCKcZ1i5gDKgWGiMJM5Sd8DmJ5rqQQerO3eNG1ml93ziW6JcMXSdCGWVjSMEGz4FK2q5ZFZRM9GXklxE2osyDZT5Wa0PKegwndvjkOsqU2EMfcnHppOXkeOgV2W4uTDNynx58ecXsWHuMnF6XxFCHzVF3OFhuLu+31vONPQCRLvWvQY5hvdzcokzBAYIkkshGHB3RZvTEg3X19aKymc+9noXh3GMVO1H45MnKBICfy7bjRMhUIug0t/ugKu/AHaBT3gd+nsR3MFa5JYlVwmB48eP3DcZAbDZfvwIDAQABAoIBAC5LsL+AvvrzBALnwokgGy1ooPw4B9JM7dnG7sytrrWvW03qyKU6z+/GNolb+8VwmQjZZcXZErl54m/4k6H15gY1//O/OnZg2JzA7VlA9yDZBjHPBApRLU/vzsJz6dEgKxuAsI+E2gbwGlOyhXjR9pjlsx243vwVuU/N79HpZke6k9NonL38gs4sWMCFkeLbqgPYU/ueWMrjX8WvFqx6wFaQpOTUZQBocKJBXOFKPiiRtnms+f6p927AwaE2t0xhpBMWDtSeDVdu01MYi8W0LWfmeCKWF8f+Kwum6xscJVAzw3NWAy5j8Pe3eSoKAEts41SpN/eIYUBG2CGQMBW28QECgYEAwAyI4BPXm/Fb8APL/QWp1SC22TPYP7bCvC8lHbq3NzPkJfE7MFLqZf++xf44u4MGjne2zyvS7z4GoRRA8mVka/3U7VInAahYCCpoOO0MrBbsKy4032cGVleU3cbn2EtagkHYq/Xue+ftxXzWNoVqjAeDyBavscoXcfuUaGcXA08CgYEAsJSxQ2qYz9pja9vXNtoiUO103kNd2re6saKm+IR6hw9sFwcsi3UncNaq3ifSRkm/os7JDqIEvYKGmQILcYNHGC1XLsL/5MOFzjp52E0ywYlrKqxKNHpDEgw06CoGnY99ZhRw3qJ5QABg45asanQ7JSRRxixyWzjgXaI9zqeagJECgYAx9dDRwjev2L0bFlHF0+hXingmzwbtpETKodUdA8rP8I7kk/na22cg+8QMS9NEbJSTEW+cO4FXPaKJ2vP+WwQh19nBl0KQetPT4/xS+s/2IMx1e0LD71BFu+j9PNZpfUjhY/HS4lqVH0PKwWwUqOaL4RSWQ9iW/sTUoSb9dfwVnwKBgCY5mgH5Eml0Yi9YIMecAu+356Oxu6B8Q2ruxexoaUnRedmmUOtDii0wGz8KKcfTcSuiTf3f7tzDY7W9rpJ9E6fVMNlLly+Db/TLzdKYK5Of1tYfA9VGXjyK7e0QI+x5wnCVnjVjwJGtQ0whMuO8k40fxp/6wSqkLlW2qGCTAawxAoGBAIGNYkVe7NBCCoH+mNZWSAuguPJc45B49ZyNWdaEfC9BuQrQRbIgt/mYXk2PGgnL9x8+FAwQe8cyEC8ZCsq8HyA1La4JXtHUdWVjqKqwTHlWyKSB663K9Nk72g+40W4gKtEseHSlTzAsNzek3ex29KCMWV6ge+1/GPYtjIf91JZH"
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := alipay.NewClient("2016081900288111", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		//SetAliPayPublicCertSN("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2TTrTo6Z9ZdrJRELO1/LsqpxFWPpdtWOjQmKhXeAlWf5QLTM612ClXxykk9q9Nf3pgByQNmC1ipFhmLISV5e/JgR1xenkcC5p4AU8LCT9IOJb413Qem9KlEFpUkgEV+Xqcq1LxKO+6YdZgt/3qoqa6h7RsTB52BL782TR8+qXVETmxiAXP7vIMhATulbvwZpvdKFywmFzyRt5XAVM49nD1rdPmtayfgUqkTGpZtOE1Lo80SNxAywPRQJDZ7umRqP88ipahhT5HB0e+WH7IgvYHtRtTw5VXupdnQigHWKBL7muTGan8JM3a1aOI7xBp7Aonw7nil6Qs9WnNPL3xckmwIDAQAB").
		SetNotifyUrl("https://www.gopay.ink")
	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "预创建创建订单")
	bm.Set("out_trade_no", "GZ201907301040355704")
	bm.Set("total_amount", "100")
	//创建订单
	aliRsp, err := client.TradePrecreate(bm)
	if err != nil {
		invoker.Logger.Error("error", elog.FieldErr(err))
		api.FailWithMessage("error", c)
		return
	}
	_, err = alipay.VerifySyncSign("", aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		invoker.Logger.Error("error", elog.FieldErr(err))
		api.FailWithMessage("error", c)
		return
	}
	api.OkWithData(aliRsp, c)
}
