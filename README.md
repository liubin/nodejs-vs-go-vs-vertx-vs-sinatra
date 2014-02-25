nodejs-vs-go-vs-vertx
=====================

simple test for nodejs vs go vs vertx


#pre-settings

$ sudo sysctl -w net.inet.tcp.msl=1000
net.inet.tcp.msl: 15000 -> 1000

$ sudo sysctl -w net.inet.ip.portrange.first=32768
net.inet.ip.portrange.first: 49152 -> 32768

**!!! reset it after test !!!**

refer: http://stackoverflow.com/questions/1216267/ab-program-freezes-after-lots-of-requests-why

