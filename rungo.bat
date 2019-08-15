@echo off
echo ///////////////run go file///////////////
:RecVar
set /p Var=Please Input run go file:
if "%Var%"=="" goto RecVar
go run %Var%
pause