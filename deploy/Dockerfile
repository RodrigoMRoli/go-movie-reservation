FROM nginx:alpine

RUN rm /etc/nginx/conf.d/default.conf

COPY /nginx/nginx.conf.template /etc/nginx/nginx.conf.template

COPY entrypoint.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]