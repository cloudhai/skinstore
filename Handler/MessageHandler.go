package Handler

import (
	"net/http"
	"skinstore/WxService"
	"skinstore/common"
	"skinstore/web/router"
)

func MsgGetHandler(params *router.Params,rw http.ResponseWriter) *common.WebResult {
	signature := params.Get("signature")
	timestamp := params.Get("timestamp")
	nonce := params.Get("nonce")
	echostr := params.Get("echostr")
	if WxService.CheckSignature(signature, nonce, timestamp) {
		rw.Write([]byte(echostr))
	} else {
		rw.Write([]byte{})
	}
	return nil
}

func MsgPostHandler(params *router.Params,rw http.ResponseWriter) *common.WebResult{
	msg := params.GetWxMsg()
	log.Infof("type:%s  content:%s",msg.MsgType,msg.Content)
	rw.Write([]byte(""))
	return nil
}