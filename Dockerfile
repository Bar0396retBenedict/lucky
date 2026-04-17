FROM scratch
COPY lucky /
EXPOSE 16601
# Using /data for persistent config storage
# Mount a volume here to persist configuration across container restarts
# e.g. docker run -v /my/local/path:/data lucky
VOLUME ["/data"]
WORKDIR /data
ENTRYPOINT ["/lucky"]
CMD ["-c", "/data/lucky.conf"]
