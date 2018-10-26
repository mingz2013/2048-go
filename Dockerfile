FROM centos
MAINTAINER mingz2013
COPY game-2048-go /sbin/game-2048-go
RUN chmod 755 /sbin/game-2048-go
ENTRYPOINT ["/sbin/game-2048-go"]
