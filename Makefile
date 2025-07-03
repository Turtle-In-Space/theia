COMP := bashly
VPATH := src
DOC_DIR = docs

# phony targets
.PHONY: all
all: theia mandoc markdown

# clean
clean: 
  rm -rvf $(DOC_DIR) theia

# create cli
theia: set.sh init.sh
  $(COMP) generate --upgrade

# create man page
mandoc: theia
  # TODO move cp man to correct dir
  $(COMP) render :mandoc $(DOC_DIR)

# create markdown help page
markdown: theia
  $(COMP) render :markdown $(DOC_DIR)
