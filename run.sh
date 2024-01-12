#!/bin/bash

PUBLIC_BACKEND_URL=${PUBLIC_BACKEND_URL:-http://localhost:3000}
PUBLIC_BACKEND_WSURL=${PUBLIC_BACKEND_WSURL:-${PUBLIC_BACKEND_URL/http:/ws:}}
PUBLIC_BACKEND_WSURL=${PUBLIC_BACKEND_WSURL/https:/wss:}

find frontend -type f -exec sed -i "s|http://localhost:3000|$PUBLIC_BACKEND_URL|g" {} +
find frontend -type f -exec sed -i "s|ws://localhost:3000|$PUBLIC_BACKEND_WSURL|g" {} +

cd frontend && npm run preview -- --port 4173 --host 0.0.0.0 &
cd backend && ./main &

wait -n