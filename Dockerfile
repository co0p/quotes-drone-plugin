FROM alpine
ADD quotes /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/quotes