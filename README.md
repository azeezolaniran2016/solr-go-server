#solr-server

This is a Go server which which uses solr as its datastore. It's to serve as a learning project.

### How to run

#### Prerequisites
- Go 1.14 (I beleive last 2 versions versions of Go should work, but I wrote this with Go 1.14)

- Solr 8.5 running locally on your machine. Follow [their tutorial](https://lucene.apache.org/solr/guide/8_5/solr-tutorial.html) and you should have the set up with the test collection called *techproducts*

- Start the app by in dev mode by runninng make dev/start

#### HTTP API Documentation

| S/N  | Endpoint  | Documentation  |
|----|---|---|
| 1  | `/ ` | returns hello world! |


### To Dos

[ ] Dockerize the app and use Docker compose

[ ] Cover with tests