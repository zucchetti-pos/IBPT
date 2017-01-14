FROM golang:1.7
ENV GOPATH /go
COPY src/github.com/compufour/IBPT/* /go/src/github.com/compufour/IBPT/
RUN cd /go/src/github.com/compufour/IBPT \
    && curl https://glide.sh/get | sh \
    && glide install \
    && go build -o /go/IBPT main.go

EXPOSE 8082

CMD ["/go/artigo-imasters"]
