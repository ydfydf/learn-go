package rpcdemo

import "errors"

//RPC最终被rpc client采用Service.Method的方式进行调用

//DemoService是RPC里面的Service
type DemoService struct {
}

type Args struct {
	A , B int //RPC里，像这种函数的输入参数，里面的值都要Public，因为要给别人调用
}

//Div是RPC里面的Method，Method是有一定的规范的
//RPC server里的函数对参数有要求，参数必须为两个，一个是输入参数args(输入参数只能一个，但不管是什么结构都行)，一个是输出参数*result(result值会改变，所以是指针)，返回error
//输入参数args是不是指针都可以，是指针的话直接使用原值，不是指针的话就将参数进行一份拷贝
func (DemoService) Div(args Args,result *float64) error {
	if args.B == 0 {
		return errors.New("Division by zero")
	}

	*result = float64(args.A) / float64(args.B)

	return nil
}