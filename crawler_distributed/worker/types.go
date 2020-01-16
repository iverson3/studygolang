package worker

// 序列化ParserFunc 以便其能在网络上传输
// 序列化的结果：  {"ParseCityList", nil}   {"ProfileParser", username, sex}
type SerializedParser struct {
	Name string         // 函数名
	Args interface{}    // 函数参数
}
