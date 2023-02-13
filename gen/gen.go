package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}
type Message struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
type Video struct {
	gorm.Model
	Id            int64  `json:"id,omitempty"`
	AuthorId      int64  `json:"author"` //改过
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	gorm.Model
	Id         int64  `json:"id,omitempty"`
	UserId     int64  `json:"user"` //改过
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
type MessageSendEvent struct {
	gorm.Model
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
type MessagePushEvent struct {
	gorm.Model
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

// 生成model和query目录的代码。
func main() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "F:\\textGit\\simple-demo-douyin\\dal\\query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	// Generate basic type-safe DAO API
	g.ApplyBasic(g.GenerateAllTable()...)
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
	db.AutoMigrate(&Comment{})
}
