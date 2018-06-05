package module


//type代表组件的类型
type Type string

//三种类型的参数
const (
	//代表组件的三种类型
	TYPE_DOWNLOADER Type= "downloader"
	TYPE_ANALYZER Type = "analyzer"
	TYPE_PIPELINE Type = "pipeline"
)

//代表三个合法类型的组件类型
var legalTypeLetterMap = map[Type]string{
	TYPE_DOWNLOADER:"D",
	TYPE_ANALYZER:"A",
	TYPE_PIPELINE:"P",
}

//方向映射
var legalLetterTypeMap = map[string]Type{
	"D": TYPE_DOWNLOADER,
	"A": TYPE_ANALYZER,
	"P": TYPE_PIPELINE,
}

//检查实列类型是否正确
func CheckType(moduleType Type,module Module) bool {
	if moduleType == "" || module == nil {
		return false
	}
	switch moduleType {
	case TYPE_DOWNLOADER:
		if _,ok:= module.(Downloader); ok {
			return true
		}
	case TYPE_ANALYZER:
		if _,ok := module.(Analyzer); ok {
			return true
		}
		
		
	}
}




