# Python安装


### Ubuntu安装Python2.7

```
sudo apt-get install python2.7
```

### Ubuntu安装Python3.6

```
sudo add-apt-repository ppa:jonathonf/python-3.6
```

* 如果上面报错add-apt-repository: command not found
```
apt-get update
sudo apt-get install software-properties-common python-software-properties
```


# Pip安装

### 命令

+ 安装包

```
pip install 包名
```

+ 更新包

```
pip install --upgrade 包名
```

+ 卸载包

```
pip uninstall 包名
```

### pip升级或者卸载包报

```
# 报错: Not uninstalling setuptools at /usr/lib/python2.7/dist-packages, owned by OS
# 解决: 在/usr/lib/python2.7/dist-packages目录下有包信息
rm -rf setuptools-3.3.egg-info
sudo pip uninstall setuptools
```
