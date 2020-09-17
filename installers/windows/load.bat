powershell Expand-Archive -LiteralPath raw.zip -DestinationPath .\
del raw.zip /q
rem del load.bat /q

rem move the tusk folder into the goroot (if they have one)
if "%GOROOT%" NEQ "" (
    move /Y src %GOROOT%/src/tusk
)