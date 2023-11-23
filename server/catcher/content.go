package catcher

import "time"

type Content struct {
	// 标题
	Title string `json:"title"`
	// 时间
	Time time.Time `json:"time"`
	// 简介(可选)
	Description string `json:"description"`
	// 封面图(可选)
	Cover string `json:"cover"`
	// 链接
	Link string `json:"link"  gorm:"index:idx"`
	// 发布人
	Author string `json:"author"  gorm:"index:idx"`
	// 来源
	Source string `json:"source"  gorm:"index:idx"`
	// 分类(可选)
	Category string `json:"category"`
}
