## Submit job

### Request
```
curl -X POST localhost:8080/submit
```

### Response
```
job:1688232899
```


## Get status

### Request
```
curl localhost:8080/status?jobId=job:1688232899 -w  "%{time_starttransfer}\n
```

### Response
```
job completed
4.011464
```

-s -o /dev/null -w  "%{time_starttransfer}\n
