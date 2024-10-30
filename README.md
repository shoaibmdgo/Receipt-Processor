# Receipt-Processor
The Receipt Processor is an open-source RESTful web service built in Go. It allows users to submit receipt data and calculates points based on specific receipt attributes. This service is designed for testing purposes, providing an easy way to submit and retrieve points based on any receipt input.

## Table of Contents
- Overview
  
- Features
  
- Getting Started
  
- API Endpoints
  
- Examples
- Contributing
- License

## Overview
The Receipt Processor allows you to:

- Submit any receipt data (e.g., retailer, items purchased, purchase date/time, and total amount) and receive a unique receiptID.
- Retrieve points earned based on receipt attributes, using the provided receiptID.
  
The service runs on localhost:8080 by default.

## Features
- Customizable Receipts: Submit any JSON-structured receipt, and the service will generate a unique receiptID.
- Retrieve Points: Use the receiptID to fetch calculated points.
- Simple to Test: Includes a sample payload to help you get started quickly.

## Getting Started
### Requirements
- Go (1.16 or above)
- PowerShell (or any similar command-line tool)

## Project Setup
### 1. Clone the Repository:
git clone https://github.com/shoaibmdgo/receipt-processor.git
cd receipt-processor

### 2. Install Dependencies:
go mod download

### 3. Run the Service:
go run main.go

The server will start on http://localhost:8080.

