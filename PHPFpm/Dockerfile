FROM php:8.3-fpm-alpine


RUN curl -sSLf \
        -o /usr/local/bin/install-php-extensions \
        https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions && \
    chmod +x /usr/local/bin/install-php-extensions && \
    install-php-extensions opentelemetry mysqli opcache ds

RUN echo -e "opcache.enable=1\n opcache.enable_cli=1\n opcache.jit_buffer_size=256M" > /usr/local/etc/php/conf.d/opcache.ini
