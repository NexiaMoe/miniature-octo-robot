**Helloion**

Simple command line to get IP Address and save to sqlite database on httpbin.org with Golang

To build this project, you need to install Golang first. Then, you can run this command to build this project:

Windows :
```powershell
go build -o helloion.exe
```

*nix :
```bash
go build -o helloion
```

To run this project, you can run this command:

```powershell
helloion.exe ip 
```

```
Command list :
ip          Get IP Address
getlog      Get IP from database
help        Help about any command

Flags:
  -h, --help   help for helloion
  -j, --json   output as json
```
