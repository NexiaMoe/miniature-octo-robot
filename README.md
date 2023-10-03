**Helloion**

Simple command line to get IP Address and save to sqlite database on httpbin.org with Golang

To build this project, you need to install Golang first. Then, you can run this command to build this project:
assuming you are using Windows OS


Windows :
```powershell
;$env:GOOS = "windows";$env:GOARCH = "amd64";go build -o helloion.exe
```

Linux build:
```powershell
;$env:GOOS = "linux";$env:GOARCH = "amd64";go build -o helloion
```

MacOS build:
```powershell
;$env:GOOS = "darwin";$env:GOARCH = "amd64";go build -o helloion
```


To run this project, you can run this command:

```powershell
helloion.exe ip 
```

```
Usage:
  helloion [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  getlog      Get log from database
  help        Help about any command
  ip          Print IP address from httpbin.org

Flags On IP Command:
  -s, --Second int     Get IP address periodically (default 5)
  -h, --help           help for ip
  -j, --json           Print JSON format
  -p, --periodically   Get IP address periodically
```

*How to use as a service*
Linux:
1. Build the project
2. Edit the helloion.service file to match binary path
3. Copy the helloion.service to /etc/systemd/system with root privilege or sudo
4. Run this command to enable the service
```bash
systemctl daemon-reload
```
5. Run this command to start the service
```bash
systemctl start helloion.service
```
6. Run this command to check the service status
```bash
systemctl status helloion.service
```

Windows:
1. Build the project
2. Download [NSSM](https://nssm.cc/download)
3. Put nssm.exe in the same directory as the binary
4. Open cmd/powershell as administrator
5. Run this command to create the service
```powershell
.\nssm.exe install helloion D:\Project\helloion\helloion.exe --config D:\Project\helloion ip -p -s 5
```
--config must be the full path to the binary

1. Run this command to start the service
```powershell
sc.exe start helloion
```
1. Run this command to check the service status
```powershell
sc.exe query helloion
```

