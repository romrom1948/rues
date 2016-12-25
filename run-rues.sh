#!/usr/bin/env sh

export RUES_DB="/home/gramsci/Documents/Perso/Geek/bano-data/rues.db"
export RUES_BACKEND_ADDR="localhost:8080"
export RUES_FRONTEND_ADDR="localhost:8081"
export RUES_TEMPLATE_ROOT="/home/gramsci/Documents/Perso/Geek/dev/go/src/github.com/romrom1948/rues/rues_frontend/rues.html"
RUES_LOGDIR="."

rues_backend > "$RUES_LOGDIR/rues_backend.log" 2>&1 &
rues_frontend > "$RUES_LOGDIR/rues_frontend.log" 2>&1 

