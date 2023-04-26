ifeq ($(OS), Windows_NT)
  CP ?= copy
  MV ?= move
  RM ?= del
else
  CP ?= cp
  MV ?= mv -f
  RM ?= rm -f
endif
