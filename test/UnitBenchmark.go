package UnitBenchmark

//单元测试demo
import "log"

func init() {
	log.Println("==============UnitBenchmark package init")
}

// 根据长宽获取面积
func GetArea(weight int, height int) int {
	return weight * height
}
