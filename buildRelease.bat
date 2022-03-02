@go build
@copy .\shapez-helper.exe .\release\shapez-helper.exe
@copy .\settings.ini .\release\settings.ini
@7z a -t7z .\release\build\shapez-helper-cli.7z .\release\shapez-helper.exe
@7z a -t7z .\release\build\shapez-helper-cli.7z .\release\settings.ini