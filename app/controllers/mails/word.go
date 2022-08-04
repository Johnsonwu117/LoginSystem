package mails

func Mail1() string {

	a := ("From:************@gmail.com\r\nTo:%v\r\nSubject:信箱驗證碼\r\n\r\n親愛的會員你好:\r\n\r\n感謝你註冊會員，以下為你的邀請碼:\r\n\r\n這是你的邀請碼:%08v\r\n\r\n請多多邀請別人。\r\n")
	return a
}
