#!/usr/bin/env bash
echo "updating node ===> $1"
cd $1
git status
git add .
git commit -m "updating assets: [$(date)]"
git push