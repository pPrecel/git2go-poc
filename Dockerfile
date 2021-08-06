FROM golang:1.15-alpine as builder

RUN apk add gcc libc-dev
RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community libgit2-dev

WORKDIR /app

COPY . .

RUN go build -o git2go main.go

FROM alpine

# RUN apk add gcc pkgconfig libc-dev

# INVESTIGATE if is alpine needed. 
# Because go binary need to know where we have C libraries (no need for other programs) 
# so maybe we can only copy C libraries from the builder (or create new step) 
# to the scratch container 
# (path to libraries should be included in the $PATH or $PKG_CONFIG_PATH or something like this)
RUN apk add --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community libgit2-dev

WORKDIR /app

COPY --from=builder /app/git2go .

ENTRYPOINT [ "/app/git2go" ]
