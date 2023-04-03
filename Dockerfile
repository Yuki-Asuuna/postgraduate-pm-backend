FROM golang:latest
WORKDIR /app
COPY . /app/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o ./output/
EXPOSE 8000
CMD ["./output/postgraduate-pm-backend"]