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

1. mongodb docker

```
sudo systemctl start docker.service
sudo docker run -d -p 27017:27017 --name mongodb mongo:4.2
```

if you already set it up

```
sudo docker start mongodb
```

2. redis

```
sudo systemctl start redis.service
```

# References

1. https://github.com/wangzitian0/golang-gin-starter-kit
2. https://github.com/Massad/gin-boilerplate
3. https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
