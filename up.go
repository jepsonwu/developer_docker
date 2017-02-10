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
	system    *uint
	dockerDir *string
}

func main() {
	var buildInfo BuildInfo

	buildInfo.project = flag.String("p", "", "project path")
	buildInfo.auth = flag.String("a", "", "docker auth")
	buildInfo.email = flag.String("email", "", "docker auth email")
	buildInfo.suffix = flag.String("suffix", "", "container suffix")
	buildInfo.system = flag.Uint("s", 1, "system type,[1-unix,2-windows]")
	buildInfo.dockerDir = flag.String("docker", "", "docker bin directory")
	help := flag.Bool("help", false, "usage help")
	flag.Parse()

	if (*help || flag.NFlag() == 0) {
		fmt.Println("Usage:", os.Args[0], " [project path] [docker auth]")
		flag.Usage()
		os.Exit(1)
	}

	if *buildInfo.project == "" {
		fmt.Println("You must configure project path")
		os.Exit(1)
	}

	if *buildInfo.auth == "" {
		fmt.Println("You must configure docker auth")
		os.Exit(1)
	}

	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currentPath = currentPath + string(filepath.Separator)

	//copy project
	*buildInfo.project, _ = filepath.Abs(*buildInfo.project)
	*buildInfo.project = filepath.Dir(*buildInfo.project + string(filepath.Separator))

	nginxConf := filepath.Join(*buildInfo.project, "nginx", "conf") + string(filepath.Separator)
	createDir(nginxConf)
	copyDir(filepath.Join(currentPath, "projects", "config", "nginx") + string(filepath.Separator), nginxConf)
	createDir(filepath.Join(*buildInfo.project, "nginx", "logs"))

	phpConf := filepath.Join(*buildInfo.project, "php", "conf") + string(filepath.Separator)
	createDir(phpConf)
	copyDir(filepath.Join(currentPath, "projects", "config", "php") + string(filepath.Separator), phpConf)
	createDir(filepath.Join(*buildInfo.project, "php", "logs"))

	dataConf := filepath.Join(*buildInfo.project, "data") + string(filepath.Separator)
	createDir(dataConf)
	copyDir(filepath.Join(currentPath, "projects", "data") + string(filepath.Separator), dataConf)

	//replace config
	replaceDir(currentPath, "%auth%", *buildInfo.auth)
	replaceDir(currentPath, "%email%", *buildInfo.email)
	replaceDir(currentPath, "%suffix%", *buildInfo.suffix)
	dockerComposeYml := filepath.Join(currentPath, "docker-compose.yml")
	replaceFile(dockerComposeYml, "%store_data%", filepath.Join(*buildInfo.project, "data"))
	replaceFile(dockerComposeYml, "%php_app%", filepath.Join(*buildInfo.project, "php"))
	replaceFile(dockerComposeYml, "%nginx_app%", filepath.Join(*buildInfo.project, "nginx"))

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

func replaceDir(path, search, replace string) (err error) {
	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() || strings.Index(f.Name(), "up") != -1 {
			//获取自己文件名
			return nil
		}

		err = replaceFile(path, search, replace)
		if err != nil {
			return err
		}

		return nil
	})

	return nil;
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

	dst, err := os.OpenFile(desFile, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func createDir(path string) {
	err := os.MkdirAll(path, 0644)
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
