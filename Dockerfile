FROM golang:1.7
COPY ibpt /go/ibpt
COPY ibpt.db /go/ibpt.db
RUN chmod +x  /go/ibpt
EXPOSE 8082

CMD ["/go/ibpt"]
