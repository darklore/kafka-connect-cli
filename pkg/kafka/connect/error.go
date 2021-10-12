package connect

type Error struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}
