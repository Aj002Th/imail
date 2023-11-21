package catcher

// Catcher 负责抓取
type Catcher interface {
	Catch() ([]Content, error)
}
