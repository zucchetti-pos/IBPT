FROM golang:1.7
ENV GOPATH /go
COPY src/github.com/eminetto/artigo-imasters/* /go/src/github.com/eminetto/artigo-imasters/
RUN cd /go/src/github.com/eminetto/artigo-imasters \
    && curl https://glide.sh/get | sh \
    && glide install \
    && go build -o /go/artigo-imasters main.go

EXPOSE 8082

CMD ["/go/artigo-imasters"]
