# Homebrew基础


### 安装

[Homebrew官网](https://brew.sh/)提示进行安装


### 软件安装路径

/usr/local/Cellar


### 基本命令

```bash
# 列举安装软件
brew list

# 搜索软件名
brew search 软件名

# 安装软件
brew install 软件名

# 更新软件
brew upgrade 软件名

# 卸载软件
brew uninstall 软件名

# 查看软件信息
brew info 软件名
```


### 查看依赖关系

```bash
# 查看所有软件依赖关系
brew deps --tree --installed

# 不被任何包依赖的包
brew leaves

# 依赖依赖某个包的所有包
brew uses $package --installed
```
