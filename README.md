nodejs-vs-go-vs-vertx
=====================

simple test for nodejs vs go vs vertx


#pre-settings

```
$ sudo sysctl -w net.inet.tcp.msl=1000
net.inet.tcp.msl: 15000 -> 1000
```
```
$ sudo sysctl -w net.inet.ip.portrange.first=32768
net.inet.ip.portrange.first: 49152 -> 32768
```
**!!! reset it after test !!!**

refer: http://stackoverflow.com/questions/1216267/ab-program-freezes-after-lots-of-requests-why

## Test Result:

###Case1: siege -b -c100 -t 60S 

####Nodejs:

>Transactions:		      140710 hits  
Availability:		      100.00 %  
Elapsed time:		       59.13 secs  
Data transferred:	        4.43 MB  
Response time:		        0.03 secs  
Transaction rate:	     2379.67 trans/sec  
Throughput:		        0.07 MB/sec  
Concurrency:		       82.92  
Successful transactions:      140710  
Failed transactions:	           0  
Longest transaction:	        3.12  
Shortest transaction:	        0.00  

####Vert.x:

>Transactions:		       91090 hits  
Availability:		       98.86 %  
Elapsed time:		       27.46 secs  
Data transferred:	        2.78 MB  
Response time:		        0.03 secs  
Transaction rate:	     3317.19 trans/sec  
Throughput:		        0.10 MB/sec  
Concurrency:		       95.78  
Successful transactions:       91090  
Failed transactions:	        1047  
Longest transaction:	        7.81  
Shortest transaction:	        0.00  

####Go:

>Transactions:		      128220 hits  
Availability:		      100.00 %  
Elapsed time:		       59.28 secs  
Data transferred:	        3.91 MB  
Response time:		        0.05 secs  
Transaction rate:	     2162.96 trans/sec  
Throughput:		        0.07 MB/sec  
Concurrency:		       99.88  
Successful transactions:      128220  
Failed transactions:	           0  
Longest transaction:	       14.71  
Shortest transaction:	        0.00  

####Sinatra:

>Transactions:		       93859 hits  
Availability:		      100.00 %  
Elapsed time:		       59.59 secs  
Data transferred:	        2.86 MB  
Response time:		        0.06 secs  
Transaction rate:	     1575.08 trans/sec  
Throughput:		        0.05 MB/sec  
Concurrency:		       99.74  
Successful transactions:       93859  
Failed transactions:	           0  
Longest transaction:	        0.22  
Shortest transaction:	        0.00  

####Sumary:

