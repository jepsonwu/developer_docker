package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"path/filepath"
	"io"
	"strings"
	"runtime"
	"os/exec"
)

type BuildInfo struct {
	project   *string
	auth      *string
	email     *string
	suffix    *string
	dockerDir *string
}

var buildInfo BuildInfo

func main() {
	buildInfo.project = flag.String("p", "", "project path")
	buildInfo.auth = flag.String("a", "", "docker auth")
	buildInfo.email = flag.String("email", "", "docker auth email")
	buildInfo.suffix = flag.String("suffix", "", "container suffix")
	buildInfo.dockerDir = flag.String("docker", "", "docker bin directory")
	help := flag.Bool("help", false, "usage help")
	flag.Parse()

	if (*help || flag.NFlag() == 0) {
		fmt.Println("Usage:", os.Args[0], " [project path] [docker auth]")
		flag.Usage()
		return
	}

	if *buildInfo.project == "" {
		fmt.Println("You must configure project path")
		return
	}

	if *buildInfo.auth == "" {
		fmt.Println("You must configure docker auth")
		return
	}

	currentPath, _ := os.Getwd()

	//copy project
	*buildInfo.project, _ = filepath.Abs(*buildInfo.project)
	*buildInfo.project = filepath.Dir(*buildInfo.project + string(filepath.Separator))

	copyDir(filepath.Join(currentPath, "projects") + string(filepath.Separator),
		*buildInfo.project + string(filepath.Separator))

	//replace config
	replaceBuildInfo()
	replaceComposeInfo()

	//exec docker-composer
	switch runtime.GOOS {
	case "windows":
		*buildInfo.dockerDir = filepath.Join(*buildInfo.dockerDir, "docker-compose.exe")
	default:
		*buildInfo.dockerDir = filepath.Join(*buildInfo.dockerDir, "docker-compose")
	}
	cmd := exec.Command(*buildInfo.dockerDir, "up")
	stdout, _ := cmd.StdoutPipe()
	reader := bufio.NewReader(stdout)

	err := cmd.Start()
	if err != nil {
		fmt.Println("exec docker composer failed:", err)
	}
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF || err != nil {
			break;
		}
		fmt.Println(string(line))
	}

	cmd.Wait()
	return
}

func replaceComposeInfo() (err error) {
	currentPath, _ := os.Getwd()
	dockerComposeYml := filepath.Join(currentPath, "docker-compose.yml")

	var replaceMap = make(map[string]string)
	replaceMap["%store_data%"] = filepath.Join(*buildInfo.project, "data")
	replaceMap["%php_app%"] = filepath.Join(*buildInfo.project, "php")
	replaceMap["%nginx_app%"] = filepath.Join(*buildInfo.project, "nginx")

	for search := range replaceMap {
		err = replaceFile(dockerComposeYml, search, replaceMap[search])
		if err != nil {
			return err
		}
	}

	return nil
}

func replaceBuildInfo() (err error) {
	currentPath, _ := os.Getwd()
	var replaceMap = make(map[string]string)
	replaceMap["%auth%"] = *buildInfo.auth
	replaceMap["%email%"] = *buildInfo.email
	replaceMap["%suffix%"] = *buildInfo.suffix

	err = filepath.Walk(currentPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if inArray(f.Name(), []string{"Dockerfile", "build.sh", "docker-compose.yml"}) != -1 {
			for search := range replaceMap {
				err = replaceFile(path, search, replaceMap[search])
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	return nil;
}

func inArray(target string, haystack []string) int {
	for index, val := range haystack {
		if val == target {
			return index
		}
	}

	return -1
}

//todo support provide array
func replaceFile(fileName, search, replace string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	tempFileName := filepath.Join(os.TempDir(), filepath.Base(fileName))
	tempFile, err := os.OpenFile(tempFileName, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFileName)

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(tempFile)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		_, err = writer.WriteString(strings.Replace(string(line), search, replace, -1) + "\n")
		if err != nil {
			return err
		}
	}

	writer.Flush()
	_, err = copyFile(tempFileName, fileName)
	if err != nil {
		return
	}

	return nil
}

func copyDir(srcDir, desDir string) (err error) {
	err = filepath.Walk(srcDir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if (path == srcDir) {
			return nil
		}

		desPath := strings.Replace(path, srcDir, "", 1);
		desFile := desDir + desPath
		if f.IsDir() {
			createDir(desFile)
		} else {
			_, err := copyFile(path, desFile)
			if err != nil {
				fmt.Println("copy file failed,file:", desFile)
			}
		}

		return nil
	})

	return nil;
}

func copyFile(srcFile, desFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(desFile, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0766)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func createDir(path string) {
	err := os.MkdirAll(path, 0766)
	if err != nil {
		fmt.Printf("Create directory \"%s\" error:%s", path, err)
	}
}

func reading() string {
	running := true
	reader := bufio.NewReader(os.Stdin)

	for running {
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("Input error")
		}

		result := string(line)
		if len(result) > 0 {
			return result
		}
	}

	return ""
}
