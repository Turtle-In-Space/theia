COMP := bashly
DOC_DIR = docs
SRC_DIR = src
COMMAND_DIR := $(SRC_DIR)/commands
MANPATH = /usr/local/man/man1
MD_DIR = $(DOC_DIR)/md
MAN_DIR = $(DOC_DIR)/man
VPATH := $(COMMAND_DIR):$(DOC_DIR):$(SRC_DIR)

.Phony := all
all: theia mandoc markdown

# create cli
theia: set.sh init.sh bashly.yml
	$(COMP) generate --upgrade

# create man page
mandoc: 
	$(COMP) render :mandoc $(MAN_DIR)
	sudo mkdir -p $(MANPATH)
	gzip $(MAN_DIR)/*.1
	sudo mv $(MAN_DIR)/*.gz $(MANPATH)

# create markdown help page
markdown: 
	$(COMP) render :markdown $(MD_DIR)
	sed -i 's/)/.md)/' $(MD_DIR)/index.md

# clean
.Phony := clean
clean: 
	@rm -rvf $(DOC_DIR) theia
	sudo rm -v $(MANPATH)/theia*
