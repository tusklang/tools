rem remove the tusk folder from the goroot (if they have one)
if "%GOROOT%" NEQ "" (
    del %GOROOT%/src/tusk /q
)

@rd /S /Q .\
del ..\Tusk /q