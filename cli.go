/*
 * @Title: cli.go
 * @Description: 用于执行命令行命令nvidia-smi
 * @Version: v1.0
 * @Company: Casia
 * @Author: hsj
 * @Date: 2021-01-25 16:12:51
 * @LastEditors: hsj
 * @LastEditTime: 2021-06-15 10:03:05
 */
package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// NvidiaSmi 获取nvidia-smi数据
func NvidiaSmi() []byte {
	stdout := new(bytes.Buffer)
	commandline := exec.Command("nvidia-smi")
	commandline.Stdout = stdout

	err := commandline.Start()
	if err != nil {
		return nil
	}
	err = commandline.Wait()
	if err != nil {
		return nil
	}
	return stdout.Bytes()
}

// TestSmi 测试数据
func TestSmi() []byte {
	fi, _ := os.Open("nvidia.txt")
	bytes, _ := ioutil.ReadAll(fi)
	return bytes
}

// GetHostName 获取主机名
func GetHostName() string {
	stdout := new(bytes.Buffer)
	commandline := exec.Command("cat", "/etc/hostname")
	commandline.Stdout = stdout

	err := commandline.Start()
	if err != nil {
		return ""
	}
	err = commandline.Wait()
	if err != nil {
		return ""
	}
	hostName := strings.Replace(stdout.String(), "\n", "", -1)
	return hostName
}
