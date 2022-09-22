## Tools that need to be installed
- [Docker](https://www.docker.com/)
- [Make](https://community.chocolatey.org/packages/make)
- [Go](https://go.dev/)

## Things that must be considered
***Make sure no postgres service is running in the background.**

## How to run the application?

- Open your terminal / cmd, and type the command on below:
    ```
    make runapi
    ```
- Make sure the service is running properly, as follows:
```$ docker ps
CONTAINER ID   IMAGE                                  COMMAND                  CREATED          STATUS          PORTS                      NAMES
a260dcd37bf0   maslow123/library-api-gateway:latest   "./main"                 16 minutes ago   Up 16 minutes   0.0.0.0:8000->8000/tcp     testapigateway
30d574ba99a1   maslow123/library-users:latest         "./main"                 16 minutes ago   Up 16 minutes   0.0.0.0:50051->50051/tcp   testapiuser
8b08ec14950c   postgres:latest                        "docker-entrypoint.s…"   19 minutes ago   Up 19 minutes   0.0.0.0:5433->5432/tcp     testdb
```
- Finish

## Documentation API
- You can see with command:
    ```make doc```
- Finish.

## How to run unit testing of each service?
- You just need to hit the command
    ```make test```
- Finish.

## Shut down all services
- Make sure you're on root folder, and hit the command
```$ make down```
- Finish.
