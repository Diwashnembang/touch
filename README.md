# Windows Touch
A simple Go implementation of the Linux touch command for Windows. Creates empty files quickly using concurrent operations.

## Overview

- Creates empty files like Linux's touch command
- Supports creating multiple files simultaneously
- Shows execution time
- Lightweight and fast

## Installation

1. Download touch.exe
2. Place it in your desired directory
3. Add the directory to your system PATH

## Usage
```powershell
# Create single file
touch file.txt

# Create multiple files
touch file1.txt file2.txt file3.txt

# Create with different extensions
touch doc.pdf image.jpg script.ps1
```
## Features

- Concurrent file creation
- Works across directories

## Limitations

- Only creates new files
- Does not modify timestamps of existing files
- Does not support Linux touch command options

## Notes

- Requires at least one filename as argument
- Logs errors if file creation fails
- Works in Command Prompt and PowerShel