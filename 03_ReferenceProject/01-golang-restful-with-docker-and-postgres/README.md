# Introduction
This project is refered from [here](https://github.com/fatihtotrakanli/simple-golang-restful-with-docker-and-postgres)  

# Table of Contents
1. Restful API for simple database.
2. Docker Compose for PostgreSQL.
3. Dockerfile for PostgreSQL.
4. Database access configuration inside code

# Questions
## Description

This project is a *Go language* learning project with simple RestFul services. It uses postgres db inside with docker-compose. You can compose with dockerfile or create your own postgres database without it. 

For run docker-compose, you need to write following commands. In your project folder,
```
      cd docker
      docker-compose up
```

then PostgreSQL works on 32300 Port (32300 -> 5432). You can access with database IDE (DataGrip, Intellij etc.) with configure port 32300.

If you want to conncect from your host system type the following command to terminal.
```
      psql -h localhost -p 32300 -d docker -U docker --password
```

For more information about it,

[Dockerize PostgreSQL](https://docs.docker.com/engine/examples/postgresql_service/#connecting-from-your-host-system)

## Database table configuration
```
      CREATE TABLE USERS (
        ID INT PRIMARY KEY,
        NAME TEXT NOT NULL,
        SURNAME TEXT NOT NULL,
        AGE INT NOT NULL
      );
      
      CREATE SEQUENCE public.users_id_seq NO MINVALUE NO MAXVALUE NO CYCLE;
      ALTER TABLE public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq');
      ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
```

## Database access configuration inside code 
Under config/config.go directory in the project, you will find database access configuration. You can change it with your custom configuration.
```
      DB_USER     = "docker"
      DB_PASSWORD = "docker"
      DB_NAME     = "docker"
      PORT = "32770"
```
## How can run?

First of all, you need to have *Go* in your computer. For Mac you can install with brew easily.

```
      brew install go
```

If everything is OK, you should encounter an output like this at terminal when wrote *go version*.

```
      go version                                    
      go version go1.9.2 darwin/amd64
```
You need to following library for the postgres.
```
      go get github.com/lib/pq
```
For run the project, in the project directory you need to write following command.

```
      go run main.go
```

If everything works correctly, you can start the CRUD operations with following URL.

```
      http://127.0.0.1:3000
```

## URL's and Example

List all of user (Need To Use GET method)
```
      http://127.0.0.1:3000/getAll
```
Add new User with JSON type ((Need To Use POST method))
```
      http://127.0.0.1:3000/newUser
      
      {
      	"name": "mockName",
      	"surname": "mockSurname",
      	"age": 30
      	}
```
List one user with the given Id (Need To Use GET method)
```
      http://127.0.0.1:3000/users/1
```
Update one user with the given Id (Need To Use PUT method)
```
      http://127.0.0.1:3000/users/1
```
Delete one user with the given Id (Need To Use DELETE method)
```
      http://127.0.0.1:3000/users/1
```
## Issue knowledge
### Issue 1:
when you run docker compose for this project. You get a error
```
=> ERROR [3/8] RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8
> gnupg, gnupg2 and gnupg1 do not seem to be installed, but one of them is required for this operation

```
To fix it: refer [here](https://stackoverflow.com/questions/50757647/e-gnupg-gnupg2-and-gnupg1-do-not-seem-to-be-installed-but-one-of-them-is-requ)
```
apt-get update && apt-get install -y gnupg
```
### Issue 2:
**Behavior**: Run docker, but it can't find public key of gpg server.
```
> cd docker
> docker-compose up
```

**Error**: 
```
=> ERROR [2/7] RUN  apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D                                                       20.6s 
------
 > [2/7] RUN  apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D:
#11 0.527 Executing: /tmp/tmp.1WQpOKm71q/gpg.1.sh --keyserver
#11 0.527 hkp://p80.pool.sks-keyservers.net:80
#11 0.527 --recv-keys
#11 0.527 58118E89F3A912897C070ADBF76221572C52609D
#11 0.530 gpg: requesting key 2C52609D from hkp server p80.pool.sks-keyservers.net
#11 20.55 ?: p80.pool.sks-keyservers.net: Host not found
#11 20.55 gpgkeys: HTTP fetch error 7: couldn't connect: Connection timed out
#11 20.56 gpg: no valid OpenPGP data found.
#11 20.56 gpg: Total number processed: 0
#11 20.56 gpg: keyserver communications error: keyserver unreachable
#11 20.56 gpg: keyserver communications error: public key not found
#11 20.56 gpg: keyserver receive failed: public key not found
------
failed to solve: rpc error: code = Unknown desc = executor failed running [/bin/sh -c apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D]: exit code: 2
```

**Root Cause**: 
- Public key using in project is expired.
- Public key maybe not exist.

**Solution**: Add new key for file: Dockerfile-Postgres
```
Remove: 
      # RUN  apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D


Add: 
      RUN set -ex \
        && for key in \
          9554F04D7259F04124DE6B476D5A82AC7E37093B \
          94AE36675C464D64BAFA68DD7434390BDBE9B9C5 \
          FD3A5288F042B6850C66B31F09FE44734EB7990E \
          71DCFD284A79C3B38668286BC97EC7A07EDE3FC1 \
          DD8F2338BAE7501E3DD5AC78C273792F7D83545D \
          B9AE9905FFD7803F25714661B63B535A4C206CA9 \
          C4F0DFFF4E8C1A8236409D08E73BC641CC11F4C8 \
          56730D5401028683275BD23C23EFEFE93C4CFFFE \
        ; do \
          gpg --keyserver pgp.mit.edu --recv-keys "$key" || \
          gpg --keyserver keyserver.pgp.com --recv-keys "$key" || \
          gpg --keyserver ha.pool.sks-keyservers.net --recv-keys "$key" ; \
        done
```
More details: [here](https://lifesaver.codes/answer/gpg-keyserver-timed-out)