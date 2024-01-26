PUBLIC_BACKEND_URL=${PUBLIC_BACKEND_URL:-http://localhost:3000}
PUBLIC_BACKEND_WSURL=${PUBLIC_BACKEND_WSURL:-${PUBLIC_BACKEND_URL/http:/ws:}}
PUBLIC_BACKEND_WSURL=${PUBLIC_BACKEND_WSURL/https:/wss:}

find /var/www/html -type f -exec sed -i "s|http://localhost:3000|$PUBLIC_BACKEND_URL|g" {} +
find /var/www/html -type f -exec sed -i "s|ws://localhost:3000|$PUBLIC_BACKEND_WSURL|g" {} +

nginx -g 'daemon off;'
