### conda 基本操作

* 创建虚拟环境

```bash
# conda create -y -n $env_name python=$version
conda create -y -n lee python=3.10
```


* 查看虚拟环境

```bash
conda env list
```


* 激活虚拟环境

```bash
# conda activate $env_name
conda activate lee

# 查看安装的 python 依赖包
conda list

# 退出虚拟环境
conda deactivate
```


* 删除虚拟环境

```bash
# conda remove -n $env_name
conda remove -n lee --all
```
