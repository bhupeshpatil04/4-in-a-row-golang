🎯 4 in a Row — Full Stack Project

A real-time multiplayer 4 in a Row / Connect Four game, playable 1v1 or vs a competitive bot, built with GoLang backend and React frontend. Real-time updates are handled via WebSockets, with a simple leaderboard and analytics-ready design.

🧠 Objective

Real-time multiplayer game server (GoLang)

Play against another player or a bot

Track and display leaderboard

Emit gameplay analytics (Kafka-ready)

Simple React frontend for interaction

🕹 Game Rules

7×6 grid

Players take turns dropping discs

First to connect 4 discs vertically, horizontally, or diagonally wins

If the board is full with no winner → draw

🏗 Project Structure
4-in-a-row-golang-final/
├─ backend/                # GoLang backend
│  ├─ cmd/
│  │  └─ server/
│  │     └─ main.go        # Server entrypoint
│  ├─ internal/
│  │  ├─ game/             # Game logic
│  │  └─ ws/               # WebSocket handler
│  ├─ go.mod
│  └─ go.sum
├─ frontend-react/         # React frontend
│  ├─ public/
│  │  └─ index.html
│  ├─ src/
│  │  ├─ index.js          # Main React app
│  │  ├─ App.js
│  │  └─ App.css
│  ├─ package.json
│  └─ package-lock.json
└─ README.md

⚡ Backend Setup (GoLang)

Navigate to backend folder:

cd backend


Install dependencies:

go mod tidy


Run the server:

go run cmd/server/main.go


Server will start at http://localhost:8080

WebSocket endpoint: ws://localhost:8080/ws

⚡ Frontend Setup (React)

Navigate to frontend folder:

cd frontend-react


Install dependencies:

npm install


Run the frontend:

npm start


Opens at http://localhost:3000

Connects to backend WebSocket automatically

🕹 How to Play

Open frontend in browser

Enter a username and join game

Wait for another player or bot to join

Click a column to drop your disc

Game updates in real-time

Winner displayed automatically; leaderboard updates

🧾 Features

Real-time Gameplay: WebSockets for instant updates

Competitive Bot: Blocks player wins and attempts its own win

Leaderboard: Tracks player wins

Rejoin Support: Players can reconnect within 30s if disconnected

Analytics Ready: Kafka events can track game metrics

💻 Tech Stack

Backend: GoLang, Gorilla WebSocket, PostgreSQL (optional for persistence)

Frontend: React, JavaScript, CSS

Real-time updates: WebSockets

Analytics (Bonus): Kafka-ready design

🔧 Notes

Focused on functionality over styling

Modern browser compatible

Backend must be running for frontend to function correctly
