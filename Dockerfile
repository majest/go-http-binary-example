FROM alpine:3.3
RUN mkdir -p /opt/app /tmp/data
ADD go-http-binary-example /opt/app/
VOLUME /tmp/data
EXPOSE 8085
CMD ["/opt/app/go-http-binary-example"]
