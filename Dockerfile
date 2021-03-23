FROM alpine
ADD main /bin/go/main
ENTRYPOINT ["/bin/go/main"]
