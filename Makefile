BIN_DIR=/usr/local/bin
COMMANDS=git-show-pr

all:
	@echo "usage: make [install|uninstall]"

install:
	@mkdir -p ${BIN_DIR}
	@echo "...installing ${COMMANDS} to ${BIN_DIR}"
	@install -m 0755 $(COMMANDS) $(BIN_DIR)
	@echo "done."

uninstall:
	@echo "..uninstalling ${COMMANDS} from ${BIN_DIR}"
	@test -d $(BIN_DIR) && \
	cd $(BIN_DIR) && \
	rm -f $(COMMANDS)
	@echo "done."
