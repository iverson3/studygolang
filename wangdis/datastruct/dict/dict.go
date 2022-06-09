package dict

// Consumer is used to traversal dict, if it returns false the traversal will be break
type Consumer func(key string, val interface{}) bool