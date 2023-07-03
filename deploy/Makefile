CURRENT_DIR := /app
CONFIG_FILE := $(CURRENT_DIR)/etc/config.yaml

# BUILD_DIR := $(CURRENT_DIR)/cmd/server/

# 两个编译后的可执行文件，会被放在容器中的根目录下
ROBOT_CHAT_EXE_PATH:=$(CURRENT_DIR)/robotChat

app-init: app-run
#app-init: app-migrate app-seed app-run
#
#app-migrate:
#	$(ROBOT_CHAT_CTL_EXE_PATH) database migrate -f $(CONFIG_FILE)
#
#app-seed:
#	$(ROBOT_CHAT_CTL_EXE_PATH) database seed -f $(CONFIG_FILE)

app-run:
	$(ROBOT_CHAT_EXE_PATH) -f $(CONFIG_FILE)
