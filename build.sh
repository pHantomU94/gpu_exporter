#!/bin/bash

###
 # @Title: build.sh
 # @Description: 一键生成release压缩包
 # @Version: v1.0
 # @Company: Casia
 # @Author: hsj
 # @Date: 2021-06-15 11:10:15
 # @LastEditors: hsj
 # @LastEditTime: 2021-06-15 11:25:52
### 

if [ ! -d "./release" ]; then
  mkdir release
fi

go build . && cp gpu_exporter ./build/gpu_exporter/gpu_exporter && tar -czvf ./release/gpu_exporter.tar.gz -C ./build/ gpu_exporter/