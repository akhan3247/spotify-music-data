'use client';

import UserProfile from "@/component/userprofile";
import TopTracks from "@/component/topracks";
import { useState } from 'react';
import TopArtists from "@/component/topartists";
import axios from "axios";
export default function Dashboard() {
  const [isTopTracksOpen, setIsTopTracksOpen] = useState(false);
  const [isTopArtistsOpen, setIsTopArtistsOpen] = useState(false);

  const handleLogout = async () => {
    try {
      const response = await axios.post("http://localhost:8080/api/logout", {}, {
        withCredentials: true,
      });
      const data = response.data;
      window.location.href = data.url;
    } catch (error) {
      console.error("Logout error:", error);
    }

  };





  return (
    <div className="min-h-screen bg-black text-white">
      {/* Navigation Bar */}
      <nav className="w-full bg-gray-900 p-4">
        <div className="container mx-auto flex justify-between items-center">
          <h1 className="text-2xl font-bold">Dashboard Page</h1>
          <button 
            type="button"
            onClick={handleLogout}
            className="px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded"
          >
            Logout
          </button>
        </div>
      </nav>
      
      {/* Main Content */}
      <div className="container mx-auto mt-10 px-4 max-w-4xl">
        {/* User Profile Section */}
        <div className="mb-4">
          <div className="w-full bg-gray-800 p-8 rounded-lg">
            <UserProfile />
          </div>
        </div>

        {/* Top Tracks Section */}
        <div className="mb-4">
          <button 
            onClick={() => setIsTopTracksOpen(!isTopTracksOpen)}
            className="w-full bg-gray-800 p-4 rounded-lg flex justify-between items-center"
          >
            <span className="text-xl font-semibold">Top Tracks</span>
            <span>{isTopTracksOpen ? '▼' : '▶'}</span>
          </button>
          {isTopTracksOpen && (
            <div className="mt-2">
              <div className="p-4 bg-gray-900 rounded-lg">
                <TopTracks />
              </div>
            </div>
          )}
        </div>

        {/* Top Artists Section */}
        <div className="mb-4">
          <button 
            onClick={() => setIsTopArtistsOpen(!isTopArtistsOpen)}
            className="w-full bg-gray-800 p-4 rounded-lg flex justify-between items-center"
          >
            <span className="text-xl font-semibold">Top Artists</span>
            <span>{isTopArtistsOpen ? '▼' : '▶'}</span>
          </button>
          {isTopArtistsOpen && (
            <div className="mt-2">
              <div className="p-4 bg-gray-900 rounded-lg">
                <TopArtists />
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}