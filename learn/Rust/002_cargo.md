# cargo 使用

### 基本命令

debug 模式：编译快，运行慢，没有进行编译优化
release 模式：编译慢，运行快

* 运行

```bash
# 默认是 debug 模式
cargo run

cargo run --release
```

* 编译

```bash
cargo build

cargo build --release
```

* 检验

```bash
cargo check
```


### 文件

* Cargo.toml

项目数据描述文件

```toml
[package]
name = "hello_world"    # 项目名称
version = "0.1.0"       # 项目版本
edition = "2021"        # Rust 大版本
```

* Cargo.lock

项目依赖详细清单
