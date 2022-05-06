FROM golang:1.18 AS build-env

WORKDIR /usr/src/app

COPY . .

RUN go mod vendor
RUN make build


FROM alpine


# Server environment variables
ENV DB_PORT=
ENV DB_USER=
ENV DB_PASSWORD=
ENV DB_NAME=
ENV JWT_SECRET=
ENV ENVIRONMENT=local
ENV PORT=

EXPOSE $PORT

WORKDIR /usr/src/app

COPY --from=build-env /usr/src/app/bin/app toggler-backend

CMD ["./toggler-backend"]