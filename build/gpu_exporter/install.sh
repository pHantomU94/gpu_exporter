#!/bin/bash
###
 # @Title: gpu_exporter
 # @Description:  gpu_exporter安装程序
 # @Version: v1.0
 # @Company: Casia
 # @Author: hsj
 # @Date: 2021-06-15 11:17:25
 # @LastEditors: hsj
 # @LastEditTime: 2021-06-15 16:52:15
### 

sudo cp gpu_exporter /usr/bin/gpu_exporter
sudo cp gpu_exporter.service /etc/systemd/system/gpu_exporter.service
sudo systemctl daemon-reload
sudo systemctl enable gpu_exporter.service
sudo systemctl start gpu_exporter.service
