PUBLIC_BACKEND_URL=${PUBLIC_BACKEND_URL:-http://localhost:3000}

find /usr/share/nginx/html -type f -exec sed -i "s|http://localhost:3000|$PUBLIC_BACKEND_URL|g" {} +

nginx -g 'daemon off;'
