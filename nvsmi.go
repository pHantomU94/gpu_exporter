/*
 * @Title:
 * @Description: 用来从nvidia-smi中提取各类监控信息
 * @Version:
 * @Company: Casia
 * @Author: hsj
 * @Date: 2021-01-25 14:53:31
 * @LastEditors: hsj
 * @LastEditTime: 2021-06-15 11:54:38
 */
package main

import (
	"regexp"
	"strconv"
	"strings"
)

// GPU资源统计
type GPUMetrics struct {
	Index       int     // GPU编号
	Fan         float32 // 风扇转速百分比
	Temp        int     // 温度值（摄氏度）
	PowerUsed   int     // 已经使用的功率
	PowerTotal  int     // 最大功率
	MemoryUsed  int     // 已用内存
	MemoryTotal int     // 总计内存
	GPUUtils    float32 // GPU使用率
}

// DeleteExtraSpace 去除多余空格
func DeleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "  ", " ", -1)     //替换tab为空格
	regstr := "\\s{2,}"                         //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)            //编译正则表达式
	s2 := make([]byte, len(s1))                 //定义字符数组切片
	copy(s2, s1)                                //将字符串复制到切片
	spcIndex := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spcIndex) > 0 {                     //找到适配项
		s2 = append(s2[:spcIndex[0]+1], s2[spcIndex[1]:]...) //删除多余空格
		spcIndex = reg.FindStringIndex(string(s2))           //继续在字符串中搜索
	}
	return string(s2)
}

// 解析参数
func extractParameters(line string, index int) *GPUMetrics {
	metrics := strings.Split(DeleteExtraSpace(line), " ")
	if len(metrics) < 13 {
		return nil
	}
	// Fan计算
	var Fan float32
	if metrics[1] == "N/A" {
		Fan = -1
	} else {
		Fan_value, _ := strconv.Atoi(strings.Split(metrics[1], "%")[0])

		Fan = float32(Fan_value) / 100
	}
	// 温度计算
	Temp, err := strconv.Atoi(strings.Split(metrics[2], "C")[0])
	if err != nil {
		Temp = -1
	}
	// 功率计算
	PowerUsed, err := strconv.Atoi(strings.Split(metrics[4], "W")[0])
	if err != nil {
		PowerUsed = -1
	}
	PowerTotal, err := strconv.Atoi(strings.Split(metrics[6], "W")[0])
	if err != nil {
		PowerTotal = -1
	}
	// 内存计算
	MemoryUsed, err := strconv.Atoi(strings.Split(metrics[8], "MiB")[0])
	if err != nil {
		MemoryUsed = -1
	}
	MemoryTotal, err := strconv.Atoi(strings.Split(metrics[10], "MiB")[0])
	if err != nil {
		MemoryTotal = -1
	}
	// GPU利用率计算
	GPU_value, _ := strconv.Atoi(strings.Split(metrics[12], "%")[0])
	GPUUtils := float32(GPU_value) / 100
	return &GPUMetrics{
		Index:       index,
		Fan:         Fan,
		Temp:        Temp,
		PowerUsed:   PowerUsed,
		PowerTotal:  PowerTotal,
		MemoryUsed:  MemoryUsed,
		MemoryTotal: MemoryTotal,
		GPUUtils:    GPUUtils,
	}
}

// ReadMetric 解析nvidia-smi数据
func ReadMetric() []*GPUMetrics {

	// content := NvidiaSmi()
	content := TestSmi()
	Datas := make([]*GPUMetrics, 0)
	if len(content) == 0 {
		return Datas
	}
	lines := strings.Split(string(content), "\n")
	testReg, _ := regexp.Compile("([0-9]+W / [0-9]+W)")
	// GPU编号
	index := 0
	// 逐行匹配
	for _, line := range lines {
		flag := testReg.MatchString(line)
		if flag {
			metric := extractParameters(line, index)
			index += 1
			Datas = append(Datas, metric)
		}
	}
	// 所有GPU参数集合
	return Datas
}
