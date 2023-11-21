package model

import "gorm.io/gen"

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE source = @source and author = @author and link = @link
	FilterWithSourceAuthorLink(source, author, link string) ([]gen.T, error)
}
