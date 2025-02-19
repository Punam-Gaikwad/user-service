FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD user-service /app/user-service

CMD ["./user-service"]