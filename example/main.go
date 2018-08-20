package main

import (
	"github.com/hqbobo/weibo"
	"fmt"
	"io/ioutil"
)

var clientid = ""
var clientsecret = ""
var code = ""

func token()  {
	tk, e := weibo.AccessToken(clientid, clientsecret, code, "")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("token:",tk)
	}
}

func share() {
	pic, err := ioutil.ReadFile("out.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	weibo.Share("", "他们说分享一个图片微博有点点难??,你也是这么认为么真的是http://shouboke.tv/", pic)
}

func main() {
	share()
}