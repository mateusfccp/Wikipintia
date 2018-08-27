FROM golang

ARG app_env
ENV APP_ENV $app_env

WORKDIR /go/src/wikipintia
COPY ./app .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["wikipintia"]


CMD if [ ${APP_ENV} = production ]; \
    then \
        wikipintia; \
    else \
        go get -u github.com/radovskyb/watcher/... && \
        echo "Watching directory in development mode..." && \
        watcher -startcmd -cmd="wikipintia"; \
    fi

EXPOSE 5000

