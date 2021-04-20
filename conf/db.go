package conf

type DBInfo struct {
	Host     string
	Port     int
	Database string
	User     string
	Pwd      string
}

type DBPool struct {
}

type DBConf struct {
	Info DBInfo
	Pool DBPool
}
