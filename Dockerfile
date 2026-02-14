FROM alpine:latest

WORKDIR /app
COPY gitlabci-cidevops-go-service .
EXPOSE 8082
CMD ["./gitlabci-cidevops-go-service"]
