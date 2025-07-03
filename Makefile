COMP := bashly
COMMAND_DIR := src/commands
DOC_DIR = docs
VPATH := $(COMMAND_DIR)

.Phony := all
all: theia mandoc markdown

# create cli
theia: set.sh init.sh
	bashly generate --upgrade

# create man page
mandoc: theia
	$(COMP) render :mandoc $(DOC_DIR)

# create markdown help page
markdown: theia
	$(COMP) render :markdown $(DOC_DIR)

.Phony := clean
# clean
clean: 
	@rm -rvf $(DOC_DIR) theia
