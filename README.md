# Spotify Music Data Visualizer

## Overview
The **Spotify Music Data Visualizer** is a web application that provides interactive visualizations of user listening habits. It integrates with Spotify's API to fetch user data, including top tracks, artists, playlists, and listening trends. The project is built with **Next.js** for the frontend and **Go** for the backend, ensuring a seamless and performant user experience.

## Features
- **Spotify Authentication**: Users can log in via Spotify OAuth.
- **User Dashboard**: Displays user profile, top artists, top tracks, and playlists.
- **Listening Trends**: Interactive charts showcasing listening history.
- **Performance Optimized**: Backend written in Go for fast API responses.
- **Responsive UI**: Built with Tailwind CSS for a smooth user experience.
- **Dark Mode Support**: Ensures accessibility and a modern UI experience.

## Tech Stack
### Frontend (Next.js)
- **Next.js**: React framework for SSR and static site generation.
- **TypeScript**: Enhances code reliability and maintainability.
- **Tailwind CSS**: Used for styling and responsiveness.
- **Recharts**: Used for interactive data visualizations.

### Backend (Go)
- **Golang**: High-performance API development.
- **Gin Framework**: Lightweight web framework for handling API requests.
- **Spotify Web API**: Fetching user listening data.
- **PostgreSQL**: Storing and managing user data.

## Installation & Setup
### Prerequisites
- Node.js & npm
- Go (latest version)
- PostgreSQL database
- Spotify Developer Account (for API credentials)

### Backend Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/spotify-music-visualizer.git
   cd spotify-music-visualizer/backend
   ```
2. Set up environment variables:
   ```env
   SPOTIFY_CLIENT_ID=your_spotify_client_id
   SPOTIFY_CLIENT_SECRET=your_spotify_client_secret
   DATABASE_URL=your_database_url
   ```
3. Install dependencies and run the backend:
   ```bash
   go mod tidy
   go run main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd ../frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Set up environment variables in `.env.local`:
   ```env
   NEXT_PUBLIC_SPOTIFY_CLIENT_ID=your_spotify_client_id
   NEXT_PUBLIC_API_URL=http://localhost:5000
   ```
4. Run the frontend:
   ```bash
   npm run dev
   ```

## Usage
- Visit `http://localhost:3000` to access the application.
- Log in with your Spotify account.
- Explore your top tracks, artists, and playlists with interactive charts.



