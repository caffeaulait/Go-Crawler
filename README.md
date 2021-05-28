# Go_Crawler



**pull elasticsearch image**

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
```

**start elasticsearch server**

```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2
```

**pull redis image**
```
docker pull redis:latest
```

**start redis server**
```
docker run -it --name redis-go -p 6379:6379 redis
```

**start worker node**

```
go run crawler_distributed/worker/server/server.go --port 9000
```

**start item saver node**

```
go run crawler_distributed/persist/server/server.go --port 1234
```

**start engine node**

```
go run crawler_distributed/main.go --itemsaver=":1234" --workers=":9000,:9001"
```

