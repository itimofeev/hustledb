FROM alpine:3.4

RUN apk add --update ca-certificates # Certificates for SSL

# Add the binary
COPY hustledb /bin/hustledb

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the axxoncloud command by default when the container starts.
CMD ["/bin/hustledb"]
