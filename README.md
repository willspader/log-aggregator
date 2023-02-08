### Mongodb

# WIP

```
docker volume create mongodata

docker run -d -p 27017:27017 --name mongo \
        -v mongodbdata:/data/db \
        -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
        -e MONGO_INITDB_ROOT_PASSWORD=secret \
        mongo:5.0
```

- You can specified the database name `log_aggregator_db` as environment variable that is used in the app

### Technical Solution

- In a real-world implementation, a sidecar would be using the service name to
