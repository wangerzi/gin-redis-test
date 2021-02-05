package redis

type UriCheckParams struct {
	Addr string `uri:"addr" binding:"required"`
	Password string `uri:"password"`
	Db int `uri:"db"`
}
type Information struct {
	Keys []string
	MemoryInfo map[string]string
	ServerInfo map[string]string
}
