# start with go
FROM golang:1.22.1-alpine AS build

# install node & make for build steps
#   (n requires bash)
RUN apk add --update npm make bash 
RUN npm install npm@latest -g 
RUN npm install n -g
RUN n lts


# workdir
RUN mkdir -p /app
WORKDIR /app

# copy files
COPY . .
RUN make clean

# install stuff
RUN make install

# build binary
RUN make build

# copy to slim container
FROM alpine:latest
COPY --from=build /app/main /
CMD ["/main"]
