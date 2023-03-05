package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type Message struct {
	gorm.Model
	Id          int64     `json:"id,omitempty"`
	Content     string    `json:"content,omitempty"`
	FromUserID  int64     `json:"from_user_id"`
	ToUserID    int64     `json:"to_user_id"`
	MessageTime time.Time `json:"column:message_time;default:CURRENT_TIMESTAMP"`
}

type Video struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	AuthorId      int64  `json:"author"` //改过
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	Title         string `json:"title"`
}
type Favorite struct {
	gorm.Model
	Id         int64 `json:"id,omitempty"`
	UserId     int64 `json:"user_id"` //改过
	VideoId    int64 `json:"video_id"`
	ActionType int32 `json:"action_type"`
}

type Comment struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	UserId     int64  `json:"user_id"` //改过
	Content    string `json:"content,omitempty"`
	VideoId    int64  `json:"video_id"`
	CreateDate string `json:"create_date"`
}

type Follow struct {
	gorm.Model
	Id         int64 `json:"id,omitempty"`
	UserId     int64 `json:"user_id"` //改过
	ToUserId   int64 `json:"to_user_id"`
	ActionType int32 `json:"action_type,omitempty"`
}

// 生成model和query目录的代码。
func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "F:\\textGit\\douyin\\dal\\query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	// Generate basic type-safe DAO API
	// 指定一个表
	user := g.GenerateModel("users")
	//videos := g.GenerateModel("videos")
	//comments := g.GenerateModel("comments")
	follows := g.GenerateModel("follows")
	//messages := g.GenerateModel("messages")
	g.ApplyBasic(user, follows)
	g.Execute()
}

// 将User结构复制给mysql数据库。
func main1() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema//
	db.AutoMigrate(&Follow{})
}
