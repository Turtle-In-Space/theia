COMP := bashly
COMMAND_DIR := src/commands
DOC_DIR = docs
VPATH := $(COMMAND_DIR)

.Phony := all
all: theia mandoc markdown

# create cli
theia: set.sh init.sh 
	$(COMP) generate --upgrade

# create man page
mandoc: theia
	$(COMP) render :mandoc $(DOC_DIR)/man

# create markdown help page
markdown: theia
	$(COMP) render :markdown $(DOC_DIR)/md
	sed -i 's/)/.md)/' docs/md/index.md

.Phony := clean
# clean
clean: 
	@rm -rvf $(DOC_DIR) theia
