@echo off

go test

:: Check if tests passed
if %errorlevel% equ 0 (
  :: If tests passed, run your main program
  echo Test: PASSED running main.go
  go run main.go deck.go
  pause
) else (
  echo Tests failed. Exiting...
  pause
)