#!/bin/bash

BASE_URL="http://localhost:8080/api"

echo "1. Create Category (Electronics)"
curl -v -X POST "$BASE_URL/categories" \
     -H "Content-Type: application/json" \
     -d '{
           "name": "Electronics",
           "description": "Gadgets and devices"
         }'
echo -e "\n\n"

echo "2. Create Product (Laptop - Linked to Category 1)"
curl -v -X POST "$BASE_URL/products" \
     -H "Content-Type: application/json" \
     -d '{
           "name": "Laptop Gaming",
           "description": "High Performance",
           "price": 25000000,
           "stock": 5,
           "category_id": 1
         }'
echo -e "\n\n"

echo "3. Get All Products (Should show category_name)"
curl -v "$BASE_URL/products"
echo -e "\n\n"

echo "4. Get Product By ID (1)"
curl -v "$BASE_URL/products/1"
echo -e "\n\n"
