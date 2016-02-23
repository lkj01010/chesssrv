package fw

import (
	"github.com/Sirupsen/logrus"
	"fmt"
)

var Log = logrus.New()

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

func PrintType(v interface{}, args...interface{}){
	if len(args) > 0 {
		fmt.Printf("Type of %s is %T\n", args[0], v)
	}else{
		fmt.Printf("Value type is %T\n", v)
	}
}

//todo
func LogInfo()
