package printSolution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionName(i interface{}) string {
	sl := strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "/")
	s := sl[len(sl)-1]
	return s
}

func PrintSolution[T any](functionName string, args ...T) string {
	s := fmt.Sprintf("%s\n\n%v\n\n--------------------------------------------------", functionName, args)
	return s
}
