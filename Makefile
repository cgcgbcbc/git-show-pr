BIN_DIR=/usr/local/bin
COMMANDS=git-show-pr

all:
	@echo "usage: make [install|uninstall]"

install:
	install -d -m 0755 $(BIN_DIR)
	install -m 0755 $(COMMANDS) $(BIN_DIR)

uninstall:
	test -d $(BIN_DIR) && \
	cd $(BIN_DIR) && \
	rm -f $(COMMANDS)
