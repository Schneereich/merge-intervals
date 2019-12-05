FROM alpine:latest
ADD merge-intervals /
ENTRYPOINT ["/merge-intervals"]