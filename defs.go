package weibo

const (
	access_token_rul = "https://api.weibo.com/oauth2/access_token?client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s"
	share_url = "https://api.weibo.com/2/statuses/share.json?access_token=%s&status=%s"
	share_pic_url = "https://api.weibo.com/2/statuses/share.json"
)

type errormsg struct {
	Error             string
	Error_code        int
	Request           string
	Error_uri         string
	Error_description string
}

type AccessTokenRsp struct {
	Access_token string
	Expires_in   string
	Remind_in    string
	Uid          string
	errormsg
}
