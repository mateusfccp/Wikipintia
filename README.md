# Wikipintia

Wikipintia is a light and simple wiki system. It's objective is to provide the essential features any wiki should have, with the less overhead as possible.

Wikipintia is designid over [KISS](https://en.wikipedia.org/wiki/KISS_principle) philosophy. This means that it's design should avoid unnecessary complexity and focus in making things the right way.

## Dependencies

### Docker

Currently, Wikipintia is being developed within docker containers, so it's easier to set up and deploy. Thus, you must install docker on your system do build and run it. Docker is available in most distros repositories. It's also available for macOS and Window$. For more info, look [here](https://docs.docker.com/engine/installation/).

### Docker-compose

Besides using docker, this system also uses `docker-compose` which makes easier to share docker containers. To install `docker-compose`, have a look on their [documentation](https://docs.docker.com/compose/install/).


## Building & running

As said above, Wikipintia is being developed within docker containers. I have no intentions of making it standalone at least until it gets on beta stage.

### Environment Variables

To properly run the server, some environment variables should be set:

#### MariaDB related
* `MYSQL_ROOT_PASSWORD`: The MariaDB root password
* `MYSQL_DATABASE`: The MariaDB database name
* `MYSQL_USER`: An non-root MariaDB user that will have access to database (as set above)
* `MYSQL_PASSWORD`: The password for this user

#### Wiki related
* `APP_ENV`: `production` for production. Any other value will start the server on development mode.
* `DB_HOST`: The host of your database. By default, it should be `data`.
* `DB_USER`: The user that is going to access database. I recommend you to use the same as `MYSQL_USER`.
* `DB_PASS`: The pass of the user that is going to access database. I recommend you to use the same as `MYSQL_PASS`.
* `DB_NAME`: It must be the same as defined on `MYSQL_DATABASE`.

### Building
1. Start `docker`.
2. Run `docker-compose build` on the root directory.

### Running
To run the server: `docker-compose up`.

The Wiki server will be listening at `localhost:5000`. The database will be exposed on `localhost:8000`, so you may access it through any MySQL client.

The container will look at any modifications on source when in development environment, and re-compile the server automatically, so you don't need to re-run `docker-compose build`.
