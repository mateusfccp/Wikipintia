FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/wikipintia
WORKDIR /go/src/wikipintia

RUN go get .
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	echo "Dev-mode listening" && \
	fresh; \
	fi

EXPOSE 5000
