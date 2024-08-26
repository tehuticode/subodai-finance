#!/bin/bash

# Create main directories
mkdir -p cmd/server
mkdir -p internal/{auth,trading,account,risk}
mkdir -p pkg/{tradingview,websocket}
mkdir -p web/{templates,static/{css,js},handlers}
mkdir -p configs
mkdir -p scripts

# Create files
touch cmd/server/main.go
touch internal/auth/auth.go
touch internal/trading/{exchange.go,order.go,position.go}
touch internal/account/account.go
touch internal/risk/risk.go
touch pkg/tradingview/integration.go
touch pkg/websocket/handler.go
touch web/templates/{layout.html,index.html,trade.html}
touch web/static/css/main.css
touch web/static/js/app.js
touch web/handlers/{home.go,trade.go}
touch configs/config.yaml
touch scripts/db_init.sql
touch go.mod README.md

echo "Project structure created successfully!"
