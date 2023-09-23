# When searching for a virtual server by name, if name matches more than one of the specified variants, e.g.
# both wildcard name and regular expression match, the first matching variant will be chosen, in the following
# order of precedence:
#
# 1. exact name
# 2. longest wildcard name starting with an asterisk, e.g. “*.example.org”
# 3. longest wildcard name ending with an asterisk, e.g. “mail.*”
# 4. first matching regular expression (in order of appearance in a configuration file)

server {
  listen 80;
  server_name ocv.sivarn.com;

  location /.well-known/acme-challenge/ {
    root /var/www/certbot;
  }

  # Redirect HTTP to HTTPS
  location / {
    return 301 https://$host$request_uri;
  }
}

server {
  listen 80;
  server_name www.ocv.sivarn.com;

  location /.well-known/acme-challenge/ {
    root /var/www/certbot;
  }

  # Redirect www to non-www
  return 301 http://ocv.sivarn.com$request_uri;
}

server {
  listen 443 ssl http2;
  server_name ocv.sivarn.com;

  add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

  ssl_certificate /etc/letsencrypt/live/ocv.sivarn.com/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/ocv.sivarn.com/privkey.pem;
  include /etc/letsencrypt/options-ssl-nginx.conf;
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

  location / {
    proxy_pass http://ocv:8080/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Forwarded-Proto $scheme;
  }
}

server {
  listen 443 ssl http2;
  server_name www.ocv.sivarn.com;

  add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

  ssl_certificate /etc/letsencrypt/live/www.ocv.sivarn.com/fullchain.pem;
  ssl_certificate_key /etc/letsencrypt/live/www.ocv.sivarn.com/privkey.pem;
  include /etc/letsencrypt/options-ssl-nginx.conf;
  ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

  # Redirect www to non-www
  return 301 https://ocv.sivarn.com$request_uri;
}

# Catch-all for unrecognised requests
server {
  listen 80 default_server;
  server_name _;
  return 444;
}