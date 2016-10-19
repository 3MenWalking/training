FROM alpine:3.4
MAINTAINER Wilson Zhang <topagentwilson@gmail.com>
Add training /usr/bin/training
Add stdmsg.txt /usr/bin/stdmsg.txt
WORKDIR /usr/bin
ENTRYPOINT ["training"]

