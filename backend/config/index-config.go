package config

import (
	"backend/config/app_config"
	"backend/config/cloudinary_config"
	"backend/config/db_config"
)

func IndexConfig() {
	app_config.ConfigAPP()
	db_config.ConfigDB()
	cloudinary_config.ConfigCloudinary()
}
