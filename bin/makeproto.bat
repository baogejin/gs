@echo off
::协议文件路径, 最后不要跟“\”符号
set SOURCE_FOLDER=..\proto\protofile

::GO编译器路径
set GO_COMPILER_PATH=.\protoc.exe
::GO文件生成路径, 最后不要跟“\”符号
set GO_TARGET_PATH=..\proto

::遍历所有文件
for /f %%i in ('dir /b "%SOURCE_FOLDER%\*.proto"') do (
    echo %GO_COMPILER_PATH% --gogofaster_out=%GO_TARGET_PATH% --proto_path=%SOURCE_FOLDER% %%i
    %GO_COMPILER_PATH% --gogofaster_out=%GO_TARGET_PATH% --proto_path=%SOURCE_FOLDER% %%i
)

echo completed!

pause