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
	case TYPE_PIPELINE:
		if _,ok := module.(Pipeline); ok {
			return true
		}
	}
	return false
}

//判断组件头是否合法
func LegalType(moduleType Type) bool {
	if _,ok:= legalTypeLetterMap[moduleType]; ok {
		return true
	}
	return false
}

//GetType 获取到组件的头部
//若给定的ID不合法 则返回false
func GetType(mid MID) (bool,Type) {
	parts,err := SplitMID(mid)
	if err != nil {
		return false,""
	}
	mt,ok := legalLetterTypeMap[parts[0]]
	return ok,mt
}



































