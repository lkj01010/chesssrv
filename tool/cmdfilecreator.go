package main
import (
	"os"
	"path/filepath"
	"github.com/lkj01010/log"
	"os/exec"
)

func create() {
	//	dir, err := ioutil.TempDir("", "stringer")
	dir, _ := filepath.Abs("../com")
	log.Info(dir)
	// Create stringer in temporary directory.
	stringer := filepath.Join(dir, "stringer.exe")
	log.Info(stringer)


	root, _ := filepath.Abs("../..")
	stringersrc := filepath.Join(root, "golang.org/x/tools/cmd/stringer/stringer.go")
	err := run("go", "build", "-o", stringer, stringersrc)
	if err != nil {
		log.Fatalf("building stringer: %s", err)
	}
	//	// Read the testdata directory.
	//	fd, err := os.Open("testdata")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer fd.Close()
	//	names, err := fd.Readdirnames(-1)
	//	if err != nil {
	//		log.Fatalf("Readdirnames: %s", err)
	//	}
	//	// Generate, compile, and run the test programs.
	//	for _, name := range names {
	//		if !strings.HasSuffix(name, ".go") {
	//			log.Errorf("%s is not a Go file", name)
	//			continue
	//		}
	//		if name == "cgo.go" && !build.Default.CgoEnabled {
	//			log.Infof("cgo is no enabled for %s", name)
	//			continue
	//		}
	//		// Names are known to be ASCII and long enough.
	//		typeName := fmt.Sprintf("%c%s", name[0] + 'A' - 'a', name[1:len(name) - len(".go")])
	//		stringerCompileAndRun(dir, stringer, typeName, name)
	//	}
	typeName := "Cmd"
	name := "cmd.go"
	stringerCompileAndRun(dir, stringer, typeName, name)
}


// stringerCompileAndRun runs stringer for the named file and compiles and
// runs the target binary in directory dir. That binary will panic if the String method is incorrect.
func stringerCompileAndRun(dir, stringer, typeName, fileName string) {
	log.Infof("run: %s %s\n", fileName, typeName)
	source := filepath.Join(dir, fileName)
//	err := copy(source, filepath.Join("testdata", fileName))
//	if err != nil {
//		log.Fatalf("copying file to temporary directory: %s", err)
//	}
	stringSource := filepath.Join(dir, typeName + "_string.go")
	// Run stringer in temporary directory.
	log.Info(stringSource)
	err := run(stringer, "-type", typeName,  "-output", stringSource, source)
	if err != nil {
		log.Fatal(err)
	}
	// Run the binary in the temporary directory.
	err = run("go", "run", stringSource, source)
	if err != nil {
		log.Fatal(err)
	}
}

// copy copies the from file to the to file.
//func copy(to, from string) error {
//	toFd, err := os.Create(to)
//	if err != nil {
//		return err
//	}
//	defer toFd.Close()
//	fromFd, err := os.Open(from)
//	if err != nil {
//		return err
//	}
//	defer fromFd.Close()
//	_, err = io.Copy(toFd, fromFd)
//	return err
//}

// run runs a single command and returns an error if it does not succeed.
// os/exec should have this function, to be honest.
func run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	create()
}