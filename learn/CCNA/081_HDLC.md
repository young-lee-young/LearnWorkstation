# HDLC（High-level Data Link Control，高级数据链路控制协议）


### 基础

是由SDLC协议发展而来

每个厂家个HDLC可能不同，因此未必能够兼容


### 作用

是一种在同步链路上传输数据的二层协议


### 数据帧格式

* 传统HDLC：传统ISO HDLC只支持单协议环境

Flag + Address + Control + Data + FCS + Flag

* Cisco HDLC：支持多协议

Flag + Address + Control + Proprietary + Data + FCS + Flag


### 配置

HDLC是同步串行接口的缺省封装格式

```sh
# 启用HDLC封装
encapsulation hdcl
```
