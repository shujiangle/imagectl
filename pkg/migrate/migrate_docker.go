package migrate

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

func Command(c string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", c)
	case "linux", "darwin":
		cmd = exec.Command("bash", "-c", c)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go read(stdout, &wg)
	go read(stderr, &wg)

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Println("启动命令出错:", err)
		return
	}

	// 并发读取标准输出和标准错误输出

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		fmt.Println("命令执行出错:", err)
		return
	}
}

func read(std io.ReadCloser, wg *sync.WaitGroup) {

	defer wg.Done()
	scanner := bufio.NewScanner(std)
	for scanner.Scan() {
		fmt.Println("标准输出:", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取标准输出出错:", err)
	}
}

// 去除url的 https:// 或者http://
func ExtractIP(url string) string {
	parts := strings.Split(url, "//")
	//fmt.Println(parts)
	if len(parts) > 1 {
		return parts[1]
	}
	return parts[0]
}
