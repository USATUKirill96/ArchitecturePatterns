# Implementing Architecture patterns with Go language

## Debug in Visual Studio Code
create .vscode/launch.json with the next configuration:
```
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "program": "${workspaceFolder}/cmd",
      "args": ["runserver"],
      "envFile": "${workspaceRoot}/.env",
    }
  ]
}
```
And run the debug mode by pressing F5