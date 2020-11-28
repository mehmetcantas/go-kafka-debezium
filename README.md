# go-kafka-debezium
A simple demonstration of how to implement Debezium and Apache Kafka connection using Golang


### How to run

First, if you don't installed Docker before you can download and install Docker to your system.

https://www.docker.com/products/docker-desktop

Open command line and change directory to project location.

Execute this commands for run Kafka, Postgresql and Zookeeper

```sh
cd go-kafka-debezium
docker-compose up
```

> If you have another PostgreSQL instance on your maching open docker-compose.yml file and change PostgreSQL port. For example -> 5555:5432

This project using Kafka package from Confluent Inc. You need to get this package for prevent any exception.

```
go get github.com/confluentinc/confluent-kafka-go/kafka
```

Connect to PostgreSQL instance (created by docker-compose file).

Execute following commands for create database and table. 

```
CREATE DATABASE "DebeziumTest"
```

Select created database and run :

```
CREATE TABLE "Product"(
   "Id" SERIAL  NOT NULL
   "Name" VARCHAR(100) NOT NULL,
   "StockQuantity" INT NOT NULL
)
```

> If you want to use another table and database change connector configuration

After did all these steps come back to project and run following command to see result

```
go run main.go
```


If you want to fix some code in this project or fix README file you can create a pull request.


Happy coding 🚀