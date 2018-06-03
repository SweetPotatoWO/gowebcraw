package module


//计数类型
type Counts struct {
	 //调用计数
	 CalledCount uint64
	 //接受计数
	 AcceptedCount uint64
	 //完成计数
	 ComletedCount uint64
	 //实时处理计数
	 HandlingNumber uint64
}

//组件的摘要结构
type SummaryStruct struct {
	 ID MID `json:"id"`
	 Called    uint64      `json:"called"`
	 Accepted  uint64      `json:"accepted"`
	 Completed uint64      `json:"completed"`
	 Handling  uint64      `json:"handling"`
	 Extra     interface{} `json:"extra,omitempty"`
}

//组件的基本接口类型
//该接口必须是并发安全的
type Module interface {
	ID() MID  //返回组件的ID
	Addr() string //返回组件的网络地址
	Score() uint64 //获取到评分
	SetScore( score uint64)  //设置评分
	ScoreCalculator() CalculateScore  //获取的评分的计算器
	CalledCount() uint64 //被调用的次数
	AcceptedCount() uint64 //被接受的次数
	CompletedCount() uint64 //完成的次数
	HandlingNumber() uint64 //正在处理的次数
	Counts() Counts  //一次获取到多少的积分
	Summary() SummaryStruct //获取到组件的摘要
}

//下载器必须实现的接口
//该接口又要并发安全？？
type Downloader interface {
	Module
	Download(req *Request) (*Response,error)
}

//分析组件必须要实现的接口
//并发安全
type Analyzer interface {
	Module
	RespPrsers()
}



