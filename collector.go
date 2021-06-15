/*
 * @Title: collector.go
 * @Description: GPU资源使用情况采集器
 * @Version: v3.0
 * @Company: Casia
 * @Author: hsj
 * @Date: 2021-01-25 10:41:23
 * @LastEditors: hsj
 * @LastEditTime: 2021-06-15 09:54:41
 */
package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// GPUManager GPU信息管理器
type GPUManager struct {
	Instance             string
	UtilUsageDesc        *prometheus.Desc // 利用率采集
	TempGaugeDesc        *prometheus.Desc // 温度采集
	PowerTotalGaugeDesc  *prometheus.Desc // 总功率
	PowerUsedGaugeDesc   *prometheus.Desc // 已用功率
	MemoryTotalGaugeDesc *prometheus.Desc // 总计内存(MiB)
	MemoryUsedGaugeDesc  *prometheus.Desc // 已用内存(MiB)
	MemoryUtilUsageDesc  *prometheus.Desc // 内存使用率
	FanUtilUsageDesc     *prometheus.Desc // 风扇转速率
}

// GetGPUState 获取统计值
// FIXME:
func (g *GPUManager) GetGPUState() []*GPUMetrics {
	return ReadMetric()
}

// Describe 返回指定的统计信息描述符
func (g *GPUManager) Describe(ch chan<- *prometheus.Desc) {
	ch <- g.UtilUsageDesc
	ch <- g.TempGaugeDesc
	ch <- g.FanUtilUsageDesc
	ch <- g.MemoryTotalGaugeDesc
	ch <- g.MemoryUsedGaugeDesc
	ch <- g.MemoryUtilUsageDesc
	ch <- g.PowerTotalGaugeDesc
	ch <- g.PowerUsedGaugeDesc
}

// Collect 执行抓取数据函数并返回数据
func (g *GPUManager) Collect(ch chan<- prometheus.Metric) {
	metrics := g.GetGPUState()
	if len(metrics) == 0 {
		return
	}
	// index, err := strconv.Atoi(strings.Split(g.Instance, "-")[4])
	// if err != nil {
	// 	index = 99
	// }
	for _, metric := range metrics {
		index := metric.Index
		label := fmt.Sprintf("GPU-%d", index)
		// 风扇转速
		if metric.Fan >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.FanUtilUsageDesc,
				prometheus.GaugeValue,
				float64(metric.Fan),
				label,
			)
		}
		// 总功率
		if metric.PowerTotal >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.PowerTotalGaugeDesc,
				prometheus.GaugeValue,
				float64(metric.PowerTotal),
				label,
			)
		}
		// 已用功率
		if metric.PowerUsed >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.PowerUsedGaugeDesc,
				prometheus.GaugeValue,
				float64(metric.PowerUsed),
				label,
			)
		}
		var memoryUtil float64
		memoryUtil = -1.0
		// 总内存
		if metric.MemoryTotal > 0 {
			ch <- prometheus.MustNewConstMetric(
				g.MemoryTotalGaugeDesc,
				prometheus.GaugeValue,
				float64(metric.MemoryTotal),
				label,
			)
			memoryUtil = 0
		}
		// 已用内存
		if metric.MemoryUsed >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.MemoryUsedGaugeDesc,
				prometheus.GaugeValue,
				float64(metric.MemoryUsed),
				label,
			)
			if memoryUtil >= 0 {
				memoryUtil = float64(metric.MemoryUsed) / float64(metric.MemoryTotal)
			}
		}
		// 内存使用率
		if memoryUtil >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.MemoryUtilUsageDesc,
				prometheus.GaugeValue,
				memoryUtil,
				label,
			)
		}
		// GPU利用率
		if metric.GPUUtils >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.UtilUsageDesc,
				prometheus.GaugeValue,
				float64(metric.GPUUtils),
				label,
			)
		}
		// 温度
		if metric.Temp >= 0 {
			ch <- prometheus.MustNewConstMetric(
				g.TempGaugeDesc,
				prometheus.GaugeValue,
				float64(metric.Temp),
				label,
			)
		}
	}

}

// NewGPUManager 初始化GPUManager
func NewGPUManager(hostName string) *GPUManager {
	return &GPUManager{
		Instance: hostName,
		UtilUsageDesc: prometheus.NewDesc(
			"GPU_utilization",
			"Current utilization of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		TempGaugeDesc: prometheus.NewDesc(
			"GPU_temperature_celsius",
			"Current temperature of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		PowerTotalGaugeDesc: prometheus.NewDesc(
			"GPU_power_total_watt",
			"Total power of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		PowerUsedGaugeDesc: prometheus.NewDesc(
			"GPU_power_used_watt",
			"Current used power of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		MemoryTotalGaugeDesc: prometheus.NewDesc(
			"GPU_memory_total_MiB",
			"Total memory of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		MemoryUsedGaugeDesc: prometheus.NewDesc(
			"GPU_memory_used_MiB",
			"Current used memroy of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		MemoryUtilUsageDesc: prometheus.NewDesc(
			"GPU_memory_utilization",
			"Current memory utilization of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
		FanUtilUsageDesc: prometheus.NewDesc(
			"GPU_fan_utilization",
			"Current fan utilization of the GPU.",
			[]string{"gpu"},
			prometheus.Labels{"instance": hostName},
		),
	}
}
