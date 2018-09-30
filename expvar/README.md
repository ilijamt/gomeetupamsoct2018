expvar
=====

To run the example execute the following commands.

```bash
docker-compose up -d
export EXPVAR_HOST=http://`docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' expvar_app_1`:6000/debug/vars
watch -n 1 curl $EXPVAR_HOST -s
```

This is the code and scripts related to the article [Exposing golang metrics](https://matoski.com/article/golang-expvar-metrics/)