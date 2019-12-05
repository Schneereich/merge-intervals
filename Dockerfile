FROM scratch
ADD merge-intervals /
ENTRYPOINT ["/merge-intervals"]