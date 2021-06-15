<!--
 * @Title: readme.md
 * @Description: 使用说明
 * @Version: v1.0
 * @Company: Casia
 * @Author: hsj
 * @Date: 2021-06-15 11:28:47
 * @LastEditors: hsj
 * @LastEditTime: 2021-06-15 14:29:10
-->

# gpu_exporter安装及使用说明
gpu_exporter是为prometheus提供的gpu资源情况收集器

- [gpu_exporter安装及使用说明](#gpu_exporter安装及使用说明)
  - [说明](#说明)
  - [安装方式](#安装方式)
  - [目前可以采集的信息](#目前可以采集的信息)
    - [**显存**](#显存)
    - [**功率**](#功率)
    - [**显卡利用率**](#显卡利用率)
    - [**风扇转速**](#风扇转速)
    - [**温度**](#温度)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>
## 说明
使用`nvidia-smi`采集GPU的各类运行数据，包括：风扇状态、功率、显存使用情况以及显卡利用率
## 安装方式
下载最新版本的`release`安装包

解压并执行安装脚本
```
tar -xzvf gpu_exporter.tar.gz
cd gpu_exporter
sudo sh install.sh
```
查看运行状态
```
systemctl status gpu_eporter.service
```

## 目前可以采集的信息
以下为能采集到的各类信息及举例
### **显存**
  * **显存总大小**
  
`GPU_memory_total_MiB`，单位为MiB
```
# HELP GPU_memory_total_MiB Total memory of the GPU.
# TYPE GPU_memory_total_MiB gauge
GPU_memory_total_MiB{gpu="GPU-0",instance="gpu"} 16160
```
  * **显存已用大小**

`GPU_memory_used_MiB`，单位为MiB
```
# HELP GPU_memory_used_MiB Current used memroy of the GPU.
# TYPE GPU_memory_used_MiB gauge
GPU_memory_used_MiB{gpu="GPU-0",instance="gpu"} 0
```
  * **显存利用率**

`GPU_memory_utilization`，浮点数
```
# HELP GPU_memory_utilization Current memory utilization of the GPU.
# TYPE GPU_memory_utilization gauge
GPU_memory_utilization{gpu="GPU-0",instance="gpu"} 0
```
### **功率**
  * **功率总大小**

`GPU_power_total_watt`，单位为W
```
# HELP GPU_power_total_watt Total power of the GPU.
# TYPE GPU_power_total_watt gauge
GPU_power_total_watt{gpu="GPU-0",instance="gpu"} 300
```
  * **功率已用大小**

`GPU_power_used_watt`，单位为W
```
# HELP GPU_power_used_watt Current used power of the GPU.
# TYPE GPU_power_used_watt gauge
GPU_power_used_watt{gpu="GPU-0",instance="gpu"} 58
```
### **显卡利用率**
  * **显卡利用率**

`GPU_utilization`，浮点数
```
# HELP GPU_utilization Current utilization of the GPU.
# TYPE GPU_utilization gauge
GPU_utilization{gpu="GPU-0",instance="gpu"} 0
```
### **风扇转速**
  * **风扇转速**

`GPU_fan_utilization`，浮点数，无风扇显卡没有该指标
```
# HELP GPU_fan_utilization Current fan utilization of the GPU.
# TYPE GPU_fan_utilization gauge
GPU_fan_utilization{gpu="GPU-5",instance="gpu"} 0.6000000238418579
```
### **温度**
  * **温度**

`GPU_temperature_celsius`，单位为°C
```
# HELP GPU_temperature_celsius Current temperature of the GPU.
# TYPE GPU_temperature_celsius gauge
GPU_temperature_celsius{gpu="GPU-0",instance="gpu"} 37
```