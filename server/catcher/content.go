package catcher

// 整理和组装内容

type Content struct {
	// 标题
	Title string `json:"title" gorm:"index:idx,unique"`
	// 简介(可选)
	Description string `json:"description"`
	// 封面图(可选)
	Cover string `json:"cover"`
	// 链接
	Link string `json:"link"  gorm:"index:idx,unique"`
	// 发布人
	Author string `json:"author"  gorm:"index:idx,unique"`
	// 来源
	Source string `json:"source"  gorm:"index:idx,unique"`
	// 分类(可选)
	Category string `json:"category"`
}
