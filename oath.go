package weibo

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"errors"
)

func AccessToken(clientid, clientsecret , code ,redirect_uri string)(*AccessTokenRsp, error) {
	url := fmt.Sprintf(access_token_rul, clientid, clientsecret, code, redirect_uri)
	resp , err := http.Post(url, "application/x-www-form-urlencoded",nil )
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	r := new(AccessTokenRsp)

	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if r.Error_code != 0 {
		return nil, errors.New(r.Error_description)
	}
	return r , nil
}