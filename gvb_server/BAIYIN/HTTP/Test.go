package HTTP

import "gvb_server/BAIYIN/BY"

type Data struct {
	SendUserId int    `json:"send_user_id"`
	RevUserId  int    `json:"rev_user_id"`
	Content    string `json:"content"`
}

func PostTest() {
	http := Post(Response{
		URL: "http://127.0.0.1:8080/api/messages",
		Header: map[string]string{
			"Token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IueZvemfsyIsIm5pY2tfbmFtZSI6IiIsInJvbGUiOjEsInVzZXJfaWQiOjEsImV4cCI6MTcyNTc2NTc4NS4xMjM0MjgsImlzcyI6Inh4In0.9yMyADWyFfv2Tv2lII51UsU_dnpXIzewuRB-y8rqM9s",
		},
		Data: Data{
			SendUserId: 1,
			RevUserId:  2,
			Content:    "123",
		},
	})
	BY.Log(http)
}
