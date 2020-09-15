.PHONY: all
all:
	$(MAKE) -f ../raw/makefile
	$(WIX)bin\candle.exe -nologo setup.wxs -o setup.wixobj -ext WixUIExtension
	$(WIX)bin\light.exe -nologo setup.wixobj -o setup.msi -ext WixUIExtension

.PHONY: clean
clean:
	-del setup.msi
	-del setup.wixobj
	-del setup.wixpdb
	$(MAKE) -f ../raw/makefile clean