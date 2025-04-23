#!/bin/bash

(cd client && npm start) & (cd server && go run .)
