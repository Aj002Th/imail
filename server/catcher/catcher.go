package catcher

// Catcher 抓取、比对和存储
type Catcher interface {
	Catch() ([]Content, error)
}
