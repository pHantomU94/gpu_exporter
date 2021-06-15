/*
 * @Title: main.go
 * @Description: 主函数入口
 * @Version: v1.0
 * @Company: Casia
 * @Author: hsj
 * @Date: 2021-01-24 22:55:36
 * @LastEditors: hsj
 * @LastEditTime: 2021-06-15 10:08:51
 */
package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	instance := GetHostName()
	if instance == "" {
		return
	}
	gpuManager := NewGPUManager(instance)

	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(gpuManager)

	gather := prometheus.Gatherers{
		reg,
	}

	h := promhttp.HandlerFor(
		gather,
		promhttp.HandlerOpts{
			ErrorLog:      logrus.New(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	logrus.Infoln("Start server at :9200")
	if err := http.ListenAndServe(":9200", nil); err != nil {
		logrus.Errorf("Error occur when start server %v", err)
		return
	}
}
