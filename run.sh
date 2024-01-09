#!/bin/bash

cd frontend && npm run preview -- --port 4173 --host 0.0.0.0 &
cd backend && ./main &

wait -n