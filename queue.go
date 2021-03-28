package broker

type GetValue struct {
	Key string
}

type PutValue struct {
	Key   string
	Value string
}
