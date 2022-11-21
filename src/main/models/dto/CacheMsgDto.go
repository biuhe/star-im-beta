package dto

// CacheMsgDto 缓存消息数据对象（从缓存查询历史消息）
type CacheMsgDto struct {
	UserIdA int64
	UserIdB int64
	Start   int64
	End     int64
	BoolRev bool
}
