# netstatgo

> John Hammond | Monday, April 10, 2023

---------------

Crappy Golang code to list local listening ports and their associated processes, much like `netstat -anob` on Windows.

To receive all possible information, run as an administrator.

## Sample Output

works on all OS and ARCH...

### Windows

```sh
netstatgo_bin_windows_amd64.exe
netstatgo_bin_windows_arm64.exe

Proto   Local Address             Foreign Address           State      PID/Name
TCP     0.0.0.0:135               0.0.0.0:0                 LISTEN     1948      /svchost.exe
TCP     0.0.0.0:445               0.0.0.0:0                 LISTEN     4         /System
TCP     0.0.0.0:903               0.0.0.0:0                 LISTEN     6092      /vmware-authd.exe
TCP     0.0.0.0:913               0.0.0.0:0                 LISTEN     6092      /vmware-authd.exe
TCP     0.0.0.0:1824              0.0.0.0:0                 LISTEN     18144     /WaveLink.exe
TCP     0.0.0.0:3389              0.0.0.0:0                 LISTEN     1220      /svchost.exe
TCP     0.0.0.0:5040              0.0.0.0:0                 LISTEN     11008     /svchost.exe
TCP     0.0.0.0:7680              0.0.0.0:0                 LISTEN     10192     /svchost.exe
TCP     0.0.0.0:49664             0.0.0.0:0                 LISTEN     1632      /lsass.exe
TCP     0.0.0.0:49665             0.0.0.0:0                 LISTEN     1536      /wininit.exe
TCP     0.0.0.0:49666             0.0.0.0:0                 LISTEN     1052      /svchost.exe
TCP     0.0.0.0:49667             0.0.0.0:0                 LISTEN     2908      /svchost.exe
TCP     0.0.0.0:49668             0.0.0.0:0                 LISTEN     3104      /svchost.exe
TCP     0.0.0.0:49669             0.0.0.0:0                 LISTEN     4884      /spoolsv.exe
TCP     0.0.0.0:49671             0.0.0.0:0                 LISTEN     1612      /services.exe
TCP     127.0.0.1:28198           0.0.0.0:0                 LISTEN     18144     /WaveLink.exe
TCP     127.0.0.1:37373           0.0.0.0:0                 LISTEN     32188     /gocode.exe
TCP     127.0.0.1:49670           0.0.0.0:0                 LISTEN     5512      /dirmngr.exe
TCP     127.0.0.1:50796           0.0.0.0:0                 LISTEN     5212      /dbus-daemon.exe
TCP     127.0.0.1:50796           0.0.0.0:0                 LISTEN     5212      /dbus-daemon.exe
TCP     127.0.0.1:56993           127.0.0.1:56994           ESTABLISHED 27928     /Zoom.exe
TCP     127.0.0.1:56994           127.0.0.1:56993           ESTABLISHED 27928     /Zoom.exe
TCP     127.0.0.1:58582           127.0.0.1:58583           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58583           127.0.0.1:58582           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58586           127.0.0.1:58587           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58587           127.0.0.1:58586           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58588           127.0.0.1:58589           ESTABLISHED 11252     /vmware.exe
TCP     127.0.0.1:58589           127.0.0.1:58588           ESTABLISHED 11252     /vmware.exe
```

### Apple MAC

```sh
netstatgo_bin_darwin_arm64
netstatgo_bin_darwin_amd64

Proto   Local Address             Foreign Address           State      PID/Name  
TCP     *:49923                   :0                        LISTEN     451       /rapportd
TCP     *:49923                   :0                        LISTEN     451       /rapportd
TCP6    fe80:1a::9e9d:2e87:a65:f375:1024 fe80:1a::76cf:7b93:a347:4e2a:1024 ESTABLISHED 465       /identityservicesd
TCP6    fe80:1a::9e9d:2e87:a65:f375:1026 fe80:1a::76cf:7b93:a347:4e2a:1025 ESTABLISHED 465       /identityservicesd
TCP6    fe80:1a::9e9d:2e87:a65:f375:1026 fe80:1a::76cf:7b93:a347:4e2a:1025 ESTABLISHED 465       /identityservicesd
UDP     *:63610                   :0                                   465       /identityservicesd
UDP     *:64418                   :0                                   465       /identityservicesd
UDP     *:54144                   :0                                   488       /sharingd
TCP     *:7000                    :0                        LISTEN     512       /ControlCenter
TCP     *:7000                    :0                        LISTEN     512       /ControlCenter
TCP     *:5000                    :0                        LISTEN     512       /ControlCenter
TCP     *:5000                    :0                        LISTEN     512       /ControlCenter
TCP     127.0.0.1:8080            :0                        LISTEN     23501     /xtemplate_bin_darwin_arm64
TCP     192.168.1.2:49854         104.46.162.225:443        ESTABLISHED 28047     /Code Helper (Plugin)
TCP     192.168.1.2:50555         34.107.243.93:443         ESTABLISHED 36010     /firefox
TCP     192.168.1.2:50095         13.248.212.111:443        ESTABLISHED 68098     /Signal Helper (Renderer)
TCP     192.168.1.2:49940         76.223.92.165:443         ESTABLISHED 68098     /Signal Helper (Renderer)
TCP     192.168.1.2:54565         52.182.143.214:443        ESTABLISHED 73422     /Code Helper (Plugin)
``` 

## To do

1. DONE: mod upgrade to V4

2. DONE: add filter by PORT or PROCESS 
- go run main.go --port 8080
- go run main.go --process chrome

3. add kill by PORT or PROCESS OR PID
- go run main.go kill --port 8080
- go run main.go kill --process chrome
- go run main.go kill --pid 123475

4. add to caddy, so we have an easy way to manage things on any machine.

5. Basic HTMX GUI. The LCI is fine, but Web is useful and easy with HTMX.