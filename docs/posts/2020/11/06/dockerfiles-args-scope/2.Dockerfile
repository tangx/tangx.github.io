ARG VERSION
FROM alpine:${VERSION}
RUN echo "stage 01"

FROM alpine:${VERSION}
RUN echo "stage 02"
