package models

// 定义请求的参数结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp  注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required"`
	Captcha    string `json:"captcha" binding:"required"`
	Gender     int    `json:"gender,omitempty" binding:"oneof=1 0 -1" `
}

// ParamActivate  发送激活码邮件参数
type ParamActivate struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// ParamLogin  登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamUpdateUserInfo struct {
	OldContent string `json:"old_content" binding:"required"`
	NewContent string `json:"new_content" binding:"required"`
	Order      string `json:"order" binding:"required,oneof=email username gender password"`
	Captcha    string `json:"captcha"`
	UserName   string `json:"user_name"`
	UserId     int64  `json:"user_id,string"`
}

// ParamChunk  板块创建信息
type ParamChunk struct {
	ChunkName    string `json:"chunk_name" binding:"required"`
	Introduction string `json:"introduction" binding:"required"`
}

// VoteData  投票数据
type ParamPostVoteData struct {
	//userID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              //帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票1，反对票-1，取消图片（0）
}

type ParamCommentVoteData struct {
	//userID 从请求中获取当前的用户
	CommentID string `json:"comment_id" binding:"required"`           //帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票1，反对票-1，取消图片（0）
}

// ParamPostList  获取帖子列表query string 参数
type ParamPostList struct {
	ChunkID int64  `json:"chunk_id,string" form:"chunk_id" example:"1"` //	可以为空
	Page    int64  `json:"page" form:"page" example:"1"`                // 页码
	Size    int64  `json:"size" form:"size" example:"10"`               // 每页数量
	Order   string `json:"order" form:"order" example:"score"`          // 排序依据
	UserID  int64  `json:"user_id,string" form:"user_id" example:"1"`
}

type ParamComment struct {
	PostID  int64  `json:"post_id,string" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type ParamSubComment struct {
	CommentId int64  `json:"comment_id,string" binding:"required"`
	PostID    int64  `json:"post_id,string" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

type ParamCommentList struct {
	CommentID int64  `json:"comment_id,string" form:"comment_id" example:"1"`
	PostID    int64  `json:"chunk_id,string" form:"post_id" example:"1"`                    // 可以为空
	Page      int64  `json:"page" form:"page" example:"1"`                                  // 页码
	Size      int64  `json:"size" form:"size" example:"10"`                                 // 每页数量
	Order     string `json:"order" form:"order" binding:"oneof= time score"example:"score"` // 排序依据
	UserID    int64  `json:"user_id,string" form:"user_id" example:"1"`
}

type ParamFocusData struct {
	UserID    string `json:"user_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0"`
}

type ParamFocusList struct {
	UserID int64 `json:"user_id,string" form:"user_id" example:"1"`
	Page   int64 `json:"page" form:"page" example:"1"` // 页码
	Size   int64 `json:"size" form:"size" example:"10"`
}
type ParamHistory struct {
	PostID string `json:"post_id" binding:"required"`
}

type ParamHistoryList struct {
	UserID int64 `json:"user_id,string" form:"user_id" example:"1"`
	Page   int64 `json:"page" form:"page" example:"1"` // 页码
	Size   int64 `json:"size" form:"size" example:"10"`
}

// ParamChunkPostList  获取社区列表query string 参数
//type ParamChunkPostList struct {
//*ParamPostList
//ChunkID int64  `json:"chunk_id" form:"chunk_id"`   //	可以为空
//}
