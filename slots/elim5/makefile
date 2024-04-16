all:init-doc

# 初始化并运行后台
doc-run:init-doc run

# 初始化文档
init-doc:
	@swag init -t Seamless-Provider,Seamless-Operator,Transfer-Provider
	
# 运行后台
run:
	@go run main.go -type=backend

# 运行api
run-api:
	@go run main.go -type=api

# git排除配置文件
skip-conf:
	@git update-index --skip-worktree ./config.yaml

# git取消排除配置文件
un-skip-conf:
	@git update-index --no-skip-worktree ./config.yaml

# 生成通用proto文件
pbs-common:
	@protoc --go_out=../ ./proto/common/*.proto

# 生成游戏proto文件
pbs-game:
	@protoc --go_out=../ ./proto/game/*.proto

# 生成所有proto
pbs-all:
#	@protoc --go_out=../ ./proto/common/*.proto ./proto/game/*.proto
	@protoc --go_out=plugins=grpc:../ ./proto/common/*.proto ./proto/game/*.proto


# pprof
pprof:
	@go tool pprof http://127.0.0.1:6655/debug/pprof/profile?seconds=30  