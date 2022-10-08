FROM alpine
ADD my_demo_svr /
ADD start.sh /
ENTRYPOINT ["sh", "start.sh"]