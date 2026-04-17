FROM scratch
COPY lucky /
EXPOSE 16601
# Using /data for persistent config storage
# Mount a volume here to persist configuration across container restarts
VOLUME ["/data"]
WORKDIR /data
ENTRYPOINT ["/lucky"]
CMD ["-c", "/data/lucky.conf"]
