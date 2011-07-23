include $(GOROOT)/src/Make.inc

TARG=termon.googlecode.com/hg
GOFILES=defs.go
CGOFILES=termon.go

GC+= -I ./_obj
LD+= -L ./_obj

package: ./_obj/$(TARG).a

include $(GOROOT)/src/Make.pkg
