FROM scratch
COPY lucky /
EXPOSE 16601
# Using /data for persistent config storage
# Mount a volume here to persist configuration across container restarts
# e.g. docker run -v /my/local/path:/data lucky
# e.g. docker run -d --restart=unless-stopped -v /my/local/path:/data -p 16601:16601 lucky
VOLUME ["/data"]
WORKDIR /data
ENTRYPOINT ["/lucky"]
CMD ["-c", "/data/lucky.conf"]
