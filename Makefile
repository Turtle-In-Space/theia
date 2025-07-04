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
mandoc: 
	$(COMP) render :mandoc $(DOC_DIR)/man
	sudo mkdir -p /usr/local/man/man1
	gzip ./docs/man/*.1
	sudo mv ./docs/man/*.gz /usr/local/man/man1

# create markdown help page
markdown: 
	$(COMP) render :markdown $(DOC_DIR)/md
	sed -i 's/)/.md)/' docs/md/index.md

# clean
.Phony := clean
clean: 
	@rm -rvf $(DOC_DIR) theia
	sudo rm /usr/local/man/man1/theia*
