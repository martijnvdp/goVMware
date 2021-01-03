# goVMware



to set credentials you can use the powershell script:
```
powershell -file .\set-credz.ps1
```

## launch.js vscode
```
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/examples/virtualmachines"
            "env": {"GOVMOMI_URL":"https://vcserver","GOVMOMI_USERNAME":"admin","GOVMOMI_PASSWORD":"secretpass","GOVMOMI_INSECURE":"1"}
        }
    ]
```
## references
https://github.com/vmware/govmomi \
https://github.com/vmware/govmomi/blob/master/govc/USAGE.md
