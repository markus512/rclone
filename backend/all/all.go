package all

import (
	// Active file systems
	_ "github.com/markus512/rclone/backend/alias"
	_ "github.com/markus512/rclone/backend/amazonclouddrive"
	_ "github.com/markus512/rclone/backend/azureblob"
	_ "github.com/markus512/rclone/backend/b2"
	_ "github.com/markus512/rclone/backend/box"
	_ "github.com/markus512/rclone/backend/cache"
	_ "github.com/markus512/rclone/backend/crypt"
	_ "github.com/markus512/rclone/backend/drive"
	_ "github.com/markus512/rclone/backend/dropbox"
	_ "github.com/markus512/rclone/backend/ftp"
	_ "github.com/markus512/rclone/backend/googlecloudstorage"
	_ "github.com/markus512/rclone/backend/http"
	_ "github.com/markus512/rclone/backend/hubic"
	_ "github.com/markus512/rclone/backend/local"
	_ "github.com/markus512/rclone/backend/mega"
	_ "github.com/markus512/rclone/backend/onedrive"
	_ "github.com/markus512/rclone/backend/opendrive"
	_ "github.com/markus512/rclone/backend/pcloud"
	_ "github.com/markus512/rclone/backend/qingstor"
	_ "github.com/markus512/rclone/backend/s3"
	_ "github.com/markus512/rclone/backend/sftp"
	_ "github.com/markus512/rclone/backend/swift"
	_ "github.com/markus512/rclone/backend/webdav"
	_ "github.com/markus512/rclone/backend/yandex"
)
