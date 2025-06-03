# Web application Dockerfile

FROM alpine:latest
RUN apk add --no-cache bash

# Add your custom commands here
RUN echo "Hello, World!" > /hello.txt

# Start the web server
CMD ["bash", "-c", "echo 'Starting web server...' && exec /usr/sbin/httpd -f -p 80"]
