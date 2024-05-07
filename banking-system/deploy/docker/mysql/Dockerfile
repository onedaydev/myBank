FROM mysql:8.0

# 임시로 설정됨 추후에 docker compose 사용해서 환경 변수 참조할 것.
ENV MYSQL_DATABASE=mydatabase \ 
    MYSQL_ROOT_PASSWORD=testpasswd \
    MYSQL_USER=myuser \
    MYSQL_PASSWORD=mypassword

EXPOSE 3306

VOLUME ["/var/lib/mysql"]

CMD ["mysqld"]
