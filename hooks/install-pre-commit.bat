@CHCP 65001
@SET PWD=%~dp0

COPY /Y %PWD%pre-commit %PWD%..\.git\hooks\

@ECHO OFF
IF /I "%ERRORLEVEL%" NEQ "0" (
    @ECHO 安装pre-commit失败，请联系程序
    @PAUSE
    @EXIT
)

@ECHO 安装pre-commit成功，可以关闭控制台
@PAUSE
