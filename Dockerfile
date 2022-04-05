FROM golang:1.18 as build_drones

COPY . /drones

WORKDIR /drones
RUN cat .env
RUN CGO_ENABLED=0 go build -o drones


FROM alpine
COPY --from=build_drones /drones/drones /drones/drones
WORKDIR /drones
RUN mkdir -p storage/logs
RUN mkdir -p storage/uploads
CMD ["./drones"]