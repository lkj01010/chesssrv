package fw

import (
	"github.com/Sirupsen/logrus"
	"fmt"
)

var Log = logrus.New()
//type logger struct {
//	l *logrus.Logger
//}
//var log logger

func init() {
	//	log.Formatter = &logrus.TextFormatter{FullTimestamp:true, DisableColors:true}	// for file
	Log.Formatter = new(logrus.TextFormatter) // default
	Log.Level = logrus.DebugLevel

	// set file output
	//	file, err := os.Create("test.log")
	//	if err != nil {
	//		log.Fatalln("fail to create test.log file!")
	//	}
	//	log.Out = file
}
//
//func LogDebug(args ...interface{}) *logger {
//	log.l.Debug(args)
//	return log
//}
//
//func LogInfo(args ...interface{}) *logger{
//	log.l.Info(args)
//	return log
//}
//
//func (l* logger)LogWarn(args ...interface{}) *logger{
//	l.l.Warn(args)
//	return l
//}
//
//func (l* logger)LogError(args ...interface{}) *logger{
//	l.l.Error(args)
//	return l
//}
//
//func (l* logger)LogFatal(args ...interface{}) *logger{
//	l.l.Fatal(args)
//	return l
//}
//
//func (l *logger)WithFields(fields map[string]interface{}) *logger{
//	l.l.WithFields(logrus.Fields(fields))
//	return l
//}


// type info
func PrintType(v interface{}, args...interface{}) {
	if len(args) > 0 {
		fmt.Printf("Type of %s is %T\n", args[0], v)
	}else {
		fmt.Printf("Value type is %T\n", v)
	}
}