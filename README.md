# SnakeOnl

## Table of Contents
1. [Introduction](#introduction)
2. [Features](#features)
3. [Getting Started](#getting-started)
4. [Installation](#installation)
5. [Usage](#usage)
6. [Docker](#docker)
7. [License](#license)

## Introduction
SnakeOnl is an online multiplayer version of the classic Snake game built with a Go backend and JavaScript frontend. This project aims to provide a fun and interactive gaming experience.

## Features
- Multiplayer support
- Real-time updates
- Responsive design
- Dockerized setup for easy deployment

## Getting Started
These instructions will help you get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go 1.16 or later
- Node.js 14.x or later
- Docker (optional, for containerized setup)

### Installation

1. Clone the repository:

git clone https://github.com/nlvphu/SnakeOnl.git
cd SnakeOnl

2. Install server dependacies:

cd server
go mod download

3. Install client dependencies:

cd ../client
npm install

### Usage

1. Start the server:

cd server
go run main.go

2. Start the client:

cd client
npm start

3. Open your browser and navigate to http://localhost:3000 to play the game.

### Docker

To run the application using Docker:

1. Build the Docker image:


docker build -t snakeonl .

2. Run the Docker container:

docker run -p 3000:3000 snakeonl

Navigate to http://localhost:3000 in your browser to access the application.


### License

