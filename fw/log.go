package fw

import (
	"fmt"
)

func init() {

	// set file output
	//	file, err := os.Create("test.log")
	//	if err != nil {
	//		log.Fatalln("fail to create test.log file!")
	//	}
	//	log.Out = file
}

//func LogDebug(args ...interface{}){
//	log.Debug(args)
//}
//
//func LogInfo(args ...interface{}){
//	log.Info(args)
//}
//
//func LogWarning(args ...interface{}){
//	log.Warning(args)
//}
//
//func LogError(args ...interface{}){
//	log.Error(args)
//}
//
//func LogDebugf(args ...interface{}){
//	log.Debugf(args)
//}
//
//func LogInfof(args ...interface{}){
//	log.Infof(args)
//}
//
//func LogWarningf(args ...interface{}){
//	log.Warningf(args)
//}
//
//func LogErrorf(args ...interface{}){
//	log.Errorf(args)
//}

// type info
func PrintType(v interface{}, args...interface{}) {
	if len(args) > 0 {
		fmt.Printf("Type of %s is %T\n", args[0], v)
	}else {
		fmt.Printf("Value type is %T\n", v)
	}
}