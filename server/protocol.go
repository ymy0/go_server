package main

type Moji_Res_Header struct {
	RspTime     int64  `json:rspTime`
	ProcessCode string `json:processCode`
}

type Moji_Res_Body struct {
	Code    int    `json:code`
	Message string `json:message`
}

type MojiResp struct {
	ResponseHeader Moji_Res_Header `json:responseHeader`
	Msgbody        Moji_Res_Body   `json:msgbody`
}

type Moji_Req_Body struct {
	Type      string `json:"type"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AlertId   string `json:"alertId"`
	Name      string `json:"name"`
	Level     string `json:"level"`
	Pubtime   string `json:"pubtime"`
	PushType  int32  `json:"pushType"`
	CityName  string `json:"cityName"`
	CityId    int    `json:"cityId"`
	Md5       string `json:"md5"`
	ClearTime string `json:"clearTime"`
}

type Moji_Req_Header struct {
	AppId       string `json:"appId"`
	AppKey      string `json:"appKey,omitempty"`
	ProcessCode string `json:"processCode"`
	Sign        string `json:"sign"`
	Timestamp   string `json:"timestamp"`
}

type MojiAlertData struct {
	RequestHeader Moji_Req_Header `json:"requestHeader"`
	MsgBody       Moji_Req_Body   `json:"msgBody"`
}
