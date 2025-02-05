'use client';

import axios from "axios";

export default function LoginButton() {
  const handleLogin = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/login');
      const data = response.data;
      window.location.href = data.url;
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  return (
    <button
      onClick={handleLogin}
      className="bg-green-500 hover:bg-green-400 text-white font-bold 
                py-4 px-8 rounded-full text-lg transition-all duration-200 
                transform hover:scale-105 shadow-lg 
                flex items-center justify-center gap-2"
    >
      <span>Login with Spotify</span>
    </button>
  );
}