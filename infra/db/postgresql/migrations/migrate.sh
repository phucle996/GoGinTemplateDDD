#!/bin/bash

export DB_URL = 
migrate -database "DB_URL" -path . "@"