// this file is only needed because of
// https://github.com/microsoft/vscode-go/issues/3086#issuecomment-605647304

module foo

go 1.14

require (
	github.com/go-redis/redis/v7 v7.2.0 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
)
