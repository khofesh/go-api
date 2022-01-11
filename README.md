# simple-go-api

Example of Rest API built with gin-gonic, gorm, and mongo-go-driver

For Fun

# Table of contents:

- [Requirements](#requirements)
- [Getting started](#getting-started)
- [References](#references)

# Requirements

- Golang (duh)
- MongoDB

# Getting started

0. start docker

```shell
sudo systemctl start docker.service
```

1. mongodb

```
docker run --name fdm-mongo -p 27017:27017 -d mongo
```

if you already set it up

```
docker start fdm-mongo
```

2. redis

```
docker run -d -p 6379:6379 --name fdm-redis redis
```

# References

1. https://github.com/wangzitian0/golang-gin-starter-kit
2. https://github.com/Massad/gin-boilerplate
3. https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
