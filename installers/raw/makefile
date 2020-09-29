ifeq ($(OS),Windows_NT)
	ZIPALL = powershell ..\raw\zip.ps1
	CLEAN = del
else
	ZIPALL = shell ../raw/zip.sh
	CLEAN = rm -f
endif

LEAVEDIR = cd ../../../

.PHONY: all
all: tusk_all oat_all undra_all
	-$(ZIPALL)

.PHONY: tusk_all
tusk_all:
	$(LEAVEDIR)tusk && $(MAKE) all

.PHONY: oat_all
oat_all:
	$(LEAVEDIR)oat && $(MAKE) all

.PHONY: undra_all
undra_all:
	$(LEAVEDIR)undra && $(MAKE) all

.PHONY: clean
clean: tusk_clean oat_clean undra_clean
	-$(CLEAN) raw.zip

.PHONY: tusk_clean
tusk_clean:
	-$(LEAVEDIR)tusk && $(MAKE) clean

.PHONY: oat_clean
oat_clean:
	-$(LEAVEDIR)oat && $(MAKE) clean

.PHONY: undra_clean
undra_clean:
	-$(LEAVEDIR)undra && $(MAKE) clean
