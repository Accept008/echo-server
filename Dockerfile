# build阶段
FROM golang:1.12-alpine AS build

# Setting the working directory of our application
WORKDIR /go/src/echo-server

# Adding all the files required to compile the application
ADD . .

# Compiling a static go application (include C libraries in the built binary)
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# Set the startup command to be our application
CMD ["./cmd"]


# 生产阶段(实现大幅度缩小镜像大小)
FROM scratch AS prod

# 从buil阶段拷贝二进制文件
COPY --from=build /go/src/echo-server/cmd .
CMD ["./cmd"]


# 镜像打包命令
# docker build -t 镜像名:版本号 .
#   eg: root@chen:/home/chen/go/src/go-debug# docker build -t echo-app:1.0.1 .

# 本地镜像启动命令
# docker run --rm -ti -p 本地端口:镜像服务端口 镜像名:版本号
#root@chen:/home/chen/go/src/go-debug# docker run --rm -ti -p 1324:1323 echo-app:1.0.1