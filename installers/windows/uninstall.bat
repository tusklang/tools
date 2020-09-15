rem remove the ka folder from the goroot (if they have one)
if "%GOROOT%" NEQ "" (
    del %GOROOT%/src/ka /q
)

@RD /S /Q .\
del ..\Ka /q