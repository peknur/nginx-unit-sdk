FROM nginx/unit:1.26.1-go1.17
COPY apps /apps
WORKDIR /apps/go
RUN mv /apps/go/app.go.txt /apps/go/app.go && \
    go mod init main && \
    go get -u unit.nginx.org/go && \
    go build -o /apps/go/app /apps/go/app.go
CMD ["unitd","--no-daemon","--control","0.0.0.0:8080"]
