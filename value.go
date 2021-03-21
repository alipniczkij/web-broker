package broker

type QueueValue struct {
	Value string `json:"v" binding:"required"`
}
