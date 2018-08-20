package weibo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	urle "net/url"
)

//share msg to sina weibo
func Share(accesstoken, status string, pic []byte) error {
	var err error
	var resp *http.Response

	url := fmt.Sprintf(share_url, accesstoken, urle.QueryEscape(status))
	if pic != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("pic", "")
		if err != nil {
			return err
		}
		_, err = io.Copy(part, bytes.NewReader(pic))
		writer.WriteField("access_token", accesstoken)
		writer.WriteField("status", status)
		err = writer.Close()
		if err != nil {
			return err
		}
		request, err := http.NewRequest("POST", share_pic_url, body)
		request.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err = http.DefaultClient.Do(request)

	} else {
		resp, err = http.Post(url, "application/x-www-form-urlencoded", nil)

	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("rsp:", string(body))

	r := new(AccessTokenRsp)

	err = json.Unmarshal(body, r)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if r.Error_code != 0 {
		return errors.New(r.Error_description)
	}
	return nil
}
