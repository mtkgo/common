package common

import "os"

const (
	CloudVendorAlibaba string = "alibabacloud"
	CloudVendorHuawei  string = "huaweicloud"
)

const (
	CloudEnvTest string = "test"
	CloudEnvProd string = "prod"
)

type CloudEnv struct {
	Vendor string
	Env    string
}

func GetCloudEnv() CloudEnv {
	env := CloudEnv{}
	{
		v := os.Getenv("CLOUD_VENDOR")
		env.Vendor = v
	}
	{
		v := os.Getenv("CLOUD_ENV")
		env.Env = v
	}
	return env
}

func IsCloud(env *CloudEnv) bool {
	return env.Vendor != ""
}

func IsCloudProd(env *CloudEnv) bool {
	return env.Env == CloudEnvProd
}
