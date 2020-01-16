package rpcdemo

import "errors"

// rpc的基本调用方式
// Service.Method

type DemoService struct {}

type Args struct {
	A, B int
}

// rpc service提供的方法只能有两个参数： args-调用该方法传递的参数 result-返回的结果(指针类型)
// 且args参数的数据类型要事先定义好； result参数必须是指针类型; 该方法本身还要返回一个错误
func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}

	*result = float64(args.A) / float64(args.B)
	return nil
}










































