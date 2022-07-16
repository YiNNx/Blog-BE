// param 包放置controller层所使用的 请求/回复 结构体
// 建议请求以Req开头，回复以Rsp开头，方便区分
// 以及，最好不要直接把model层的结构体序列化作为回应，可以降低耦合度
package param

type Stats struct {
	Likes    int `json:"likes,omitempty"`
	Views    int `json:"views,omitempty"`
	Comments int `json:"comments,omitempty"`
}

type PostOutline struct {
	Pid     string   `json:"pid,omitempty"`
	Status  int      `json:"status,omitempty"`
	Title   string   `json:"title,omitempty"`
	Time    string   `json:"time,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Excerpt string   `json:"excerpt,omitempty"`
	Stats   Stats    `json:"stats,omitempty"`
}

type ResponseGetPosts struct {
	Posts []PostOutline `json:"posts,omitempty"`
}

type ResponseLogIn struct {
	Token string `json:"token,omitempty"`
}

type ResponseGetPostByPid struct {
	Pid     string   `json:"pid,omitempty"`
	Status  int      `json:"status,omitempty"`
	Type    int      `json:"type,omitempty"`
	Title   string   `json:"title,omitempty"`
	Time    string   `json:"time,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Content string   `json:"content,omitempty"`
	Stats   Stats    `json:"stats,omitempty"`
}

type ResponseGetTags struct {
	Tags []string `json:"tags,omitempty"`
}

type RequestStatus struct {
	Status bool `json:"status,omitempty" validate:"required"`
}

type ResponseGetComments struct {
	Cid       string `json:"cid,omitempty"`
	ParentCid string `json:"parent_cid,omitempty"`
	Time      string `json:"time,omitempty"`
	From      string `json:"from,omitempty"`
	FromUrl   string `json:"from_url,omitempty"`
	Content   string `json:"content,omitempty"`
}

type RequestComment struct {
	From      string `json:"from,omitempty" validate:"required"`
	To        string `json:"to,omitempty"`
	ParentCid string `json:"parent_cid,omitempty"`
	FromUrl   string `json:"from_url,omitempty"`
	Email     string `json:"email,omitempty"`
	Content   string `json:"content,omitempty" validate:"required"`
}

type ResponseStats struct {
	Posts     int `json:"posts,omitempty"`
	Timestamp int `json:"timestamp,omitempty"`
	Views     int `json:"views,omitempty"`
	Likes     int `json:"likes,omitempty"`
	Comments  int `json:"comments,omitempty"`
}

type RequestNewPost struct {
	Status  int      `json:"status,omitempty"`
	Type    int      `json:"type,omitempty"`
	Title   string   `json:"title,omitempty" validate:"required"`
	Tags    []string `json:"tags,omitempty"`
	Excerpt string   `json:"excerpt,omitempty"`
	Content string   `json:"content,omitempty" validate:"required"`
}

type ResponseNewPost struct {
	Pid string `json:"pid,omitempty"`
}

type ResponseNewComment struct {
	Cid string `json:"cid,omitempty"`
}

type RequestUpdatePost struct {
	Status  int      `json:"status,omitempty"`
	Type    int      `json:"type,omitempty"`
	Title   string   `json:"title,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Excerpt string   `json:"excerpt,omitempty"`
	Content string   `json:"content,omitempty"`
}
