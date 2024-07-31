package async

import (
	"context"
	"gt-go/webook/internal/service/sms"
)

type SMSService struct {
	svc sms.Service
	//repo repository.SMSAysncReqRepository
}

func NewSMSService() *SMSService {
	return &SMSService{}
}

// TODO 同步转异步
// 设计一个新的容错机制，同步转异步的容错机制。当满足以下两个条件中的任何一个时，将请求转储到数据库，后续
// 再另外启动一个 goroutine 异步发送出去。
// • 触发了限流。
// • 判定服务商已经崩溃。
// 要求：
// • 如何判定服务商已经崩溃，不允许使用课程上的判断机制，你需要设计一个新的判断机制，并且解释这种判定机制
// 的决策理由。
// • 控制异步重试次数，转储到数据库之后，可以重试 N 次，重试间隔你可以自由决策。
// • 不允许写死任何参数，即用户必须可以控制控制参数。
// • 保持面向接口和依赖注入风格。
// • 写明这种容错机制适合什么场景，并且有什么优缺点。
// • 针对提出的缺点，写出后续的改进方案。
// • 提供单元测试。
func (s SMSService) StartAsync() {

}

func (s SMSService) Send(ctx context.Context, biz string, args []string, numbers ...string) error {
	err := s.svc.Send(ctx, biz, args, numbers...)
	//TODO 崩溃了咋搞
	if err != nil {
		return err
	}
	return nil
}
