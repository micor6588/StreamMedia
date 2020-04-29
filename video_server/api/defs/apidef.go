package defs

// UserCredential 用户信息结构体
type UserCredential struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// VideoInfo  上传视频信息结构体
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

// Comment 视频评论结构体
type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	Username string // login name
	TTL      int64  //检查用户是否国企
}
