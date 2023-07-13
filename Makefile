CURRENT_DIR := $(shell pwd)
CONFIG_FILE := $(CURRENT_DIR)/config.yaml

# 设定需要编译的go文件目录
BUILD_EXE_PATH := $(CURRENT_DIR)/app/main.go

PATH_BUILD := $(CURRENT_DIR)/.build
PATH_BUILD_LINUX := $(PATH_BUILD)/linux
PATH_BUILD_WINDOWS := $(PATH_BUILD)/windows
PATH_BUILD_MAC_OS := $(PATH_BUILD)/macos

# 将编译好的执行文件，放入根目录下
ROBOT_CHAT_EXE_PATH:=$(CURRENT_DIR)/robotChat

ROBOT_CHAT_EXE_BUILD_PATH_LINUX := $(PATH_BUILD_LINUX)/robotChat

ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS := $(PATH_BUILD_WINDOWS)/robotChat.exe

ROBOT_CHAT_EXE_BUILD_PATH_MAC_OS := $(PATH_BUILD_MAC_OS)/robotChat

DEPLOY_ROBOT_CHAT_EXE_PATH:=$(CURRENT_DIR)/deploy/robotChat

DEPLOY_ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS:=$(CURRENT_DIR)/deploy/robotChat.exe

# 本地windows运行路径
ROBOT_CHAT_EXE_PATH_WINDOWS := $(CURRENT_DIR)/robotChat.exe


app-init: app-run
#app-init: app-migrate app-seed app-run
#app-init-db: app-migrate app-seed

#app-migrate:
#	go build -o $(ROBOT_CHAT_CTL_EXE_PATH) $(BUILD_CTL_PATH)
#	$(ROBOT_CHAT_CTL_EXE_PATH) database migrate -f $(CONFIG_FILE)
#
#app-seed:
#	go build -o $(ROBOT_CHAT_CTL_EXE_PATH) $(BUILD_CTL_PATH)
#	$(ROBOT_CHAT_CTL_EXE_PATH) database seed -f $(CONFIG_FILE)

app-run:
	go build -o $(ROBOT_CHAT_EXE_PATH) $(BUILD_EXE_PATH)
	$(ROBOT_CHAT_EXE_PATH) -f $(CONFIG_FILE)


# ------

app-build-linux:
	CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o $(ROBOT_CHAT_EXE_BUILD_PATH_LINUX) $(BUILD_EXE_PATH)
	cp $(ROBOT_CHAT_EXE_BUILD_PATH_LINUX) $(DEPLOY_ROBOT_CHAT_EXE_PATH)

app-build-windows:
	set CGO_ENABLED=0 && set GOOS=windows && set GOARCH=amd64 && go build -o $(ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS) $(BUILD_EXE_PATH)
	copy $(ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS) $(ROBOT_CHAT_EXE_PATH_WINDOWS)

app-build-windows-power-shell:
	$Env:CGO_ENABLED=0; $Env:GOOS="windows"; $Env:GOARCH="amd64"; go build -o $(ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS) $(BUILD_EXE_PATH)
	Copy-Item $(ROBOT_CHAT_EXE_BUILD_PATH_WINDOWS) $(ROBOT_CHAT_EXE_PATH_WINDOWS)

app-build-macos:
	CGO_ENABLED=0  GOOS=darwin  GOARCH=arm64 go build -o $(ROBOT_CHAT_EXE_BUILD_PATH_MAC_OS) $(BUILD_EXE_PATH)
	cp $(ROBOT_CHAT_EXE_BUILD_PATH_MAC_OS) $(DEPLOY_ROBOT_CHAT_EXE_PATH)


