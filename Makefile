# vim: tabstop=4:softtabstop=4:shiftwidth=4:noexpandtab

THISPROJECT="licensegen"

SRCPATH=.
BUILDDIR=.
INSTALLDIR=~/scripts
GO_BIN=go

ALLTARGETS=licensegen

####################################################################

all: $(ALLTARGETS)
	@echo "* Compiled all targets for $(THISPROJECT)."
	@exit 0

$(ALLTARGETS): init 
	@$(GO_BIN) build -o $(BUILDDIR)/$@ $(SRCPATH)/$@.go
	@echo "* Build finished: '$@'"
	@exit 0

strip:
	@$(foreach t,$(ALLTARGETS),if [ -f $(BUILDDIR)/$(t) ]; then strip --strip-all $(BUILDDIR)/$(t) ; fi; )
	@echo "* Stripped all built files."
	@exit 0

test:
	@$(GO_BIN) test -v
	@echo "* Testing complete."
	@exit 0

install:
	@$(foreach t,$(ALLTARGETS),if [ -f $(BUILDDIR)/$(t) ]; then /bin/cp $(BUILDDIR)/$(t) $(INSTALLDIR)/$(t) ; fi; )
	@echo "* Installed all built files."
	@exit 0

init:
	@if [ ! -e $(BUILDDIR) ]; then mkdir $(BUILDDIR); fi
	@exit 0

clean: 
	@rm $(addprefix $(BUILDDIR)/,$(ALLTARGETS))
	@echo "* Cleaned all built files."
	@exit 0
