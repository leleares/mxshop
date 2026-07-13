### 项目依赖初始化
1. clone 项目到本地
2. 执行 go mod download 会尽量按照项目已有的 go.mod / go.sum 下载依赖，不主动清理或改写依赖声明。
3. 执行 go mod tidy 会根据当前源码重新计算依赖，可能修改 go.mod 和 go.sum。