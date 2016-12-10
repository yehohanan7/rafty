# rafty
A sample application which uses raft to maintain state

## overview
The application exposes the below http endpoints to maintain a calculator across 3 distributed nodes/peers

### starting node1 on port 8001
```bash
go run main.go 1
```

### starting node2 on port 8002
```bash
go run main.go 2
```

### starting node3 on port 8003
```bash
go run main.go 3
```

### add a number on node1 and get it from node2
```bash
curl -X POST http://localhost:8001/add -d '{"X": 5}'

curl http://localhost:8002/value
```


