FROM golang:1.13.7
RUN mkdir /Espresso
Add . /Espresso
WORKDIR /Espresso
RUN go build -o main main.go
EXPOSE 8080
CMD ["/Espresso/main"]
