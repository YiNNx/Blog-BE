// param 包放置controller层所使用的 请求/回复 结构体
// 建议请求以Req开头，回复以Rsp开头，方便区分
// 以及，最好不要直接把model层的结构体序列化作为回应，可以降低耦合度
package param

type Stats struct {
	Likes    int `json:"likes"`
	Views    int `json:"views"`
	Comments int `json:"comments"`
}

type PostOutline struct {
	Pid       string   `json:"pid"`
	Status    int      `json:"status"`
	Title     string   `json:"title"`
	Time      string   `json:"time"`
	Tags      []string `json:"tags"`
	Excerpt   string   `json:"excerpt"`
	Stats     Stats    `json:"stats"`
	IsDeleted bool     `json:"is_deleted"`
}

type ResponseGetPosts struct {
	Posts []PostOutline `json:"posts"`
}

type ResponseLogIn struct {
	Token string `json:"token"`
}

type ResponseGetPostByPid struct {
	Pid     string   `json:"pid"`
	Status  int      `json:"status"`
	Type    int      `json:"type"`
	Title   string   `json:"title"`
	Time    string   `json:"time"`
	Tags    []string `json:"tags"`
	Content string   `json:"content"`
	Stats   Stats    `json:"stats"`
}

type ResponseGetTags struct {
	Tags []string `json:"tags"`
}

type RequestStatus struct {
	Status bool `json:"status" validate:"required"`
}

type ResponseGetComments struct {
	Cid       string `json:"cid"`
	ParentCid string `json:"parent_cid"`
	Time      string `json:"time"`
	From      string `json:"from"`
	FromUrl   string `json:"from_url"`
	Content   string `json:"content"`
}

type RequestComment struct {
	From      string `json:"from" validate:"required"`
	To        string `json:"to"`
	ParentCid string `json:"parent_cid"`
	FromUrl   string `json:"from_url"`
	Email     string `json:"email"`
	Content   string `json:"content" validate:"required"`
}

type ResponseStats struct {
	Posts     int `json:"posts"`
	Timestamp int `json:"timestamp"`
	Views     int `json:"views"`
	Likes     int `json:"likes"`
	Comments  int `json:"comments"`
}

type RequestNewPost struct {
	Status  int      `json:"status"`
	Type    int      `json:"type"`
	Title   string   `json:"title" validate:"required"`
	Tags    []string `json:"tags"`
	Excerpt string   `json:"excerpt"`
	Content string   `json:"content" validate:"required"`
}

type ResponseNewPost struct {
	Pid string `json:"pid"`
}

type ResponseNewComment struct {
	Cid string `json:"cid"`
}

type RequestUpdatePost struct {
	Status  int      `json:"status"`
	Type    int      `json:"type"`
	Title   string   `json:"title"`
	Tags    []string `json:"tags"`
	Excerpt string   `json:"excerpt"`
	Content string   `json:"content"`
}
