#!/bin/bash +ex

echo "[get /v1/skills]"
hey -c 1000 -n 100000 -m GET http://localhost:7279/v1/skills?limit=25

echo "[post /v1/skills]"
hey -c 1000 -n 100000 -m POST http://localhost:7279/v1/skills

echo "[patch /v1/skills]"
hey -c 1000 -n 100000 -m PATCH http://localhost:7279/v1/skills

echo "[delete /v1/skills]"
hey -c 1000 -n 100000 -m DELETE http://localhost:7279/v1/skills
