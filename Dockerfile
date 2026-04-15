FROM scratch
COPY lucky /
EXPOSE 16601
# Using /data for persistent config storage
WORKDIR /data
ENTRYPOINT ["/lucky"]
CMD ["-c", "/data/lucky.conf"]
