import LoginButton from "@/component/loginbutton";
import { Suspense } from "react";
export default function Home() {
  return (
    <main className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-r from-black to-green-900">
      <div className="text-center">
        <h1 className="text-6xl font-bold text-white mb-8">
          Welcome to Music Data
        </h1>
        <p className="text-xl text-gray-300 mb-8">
          Discover your Spotify listening habits
        </p>
        <div className="inline-block bg-green-900 p-4 rounded-full mt-8">
          <Suspense fallback={<div>Loading...</div>}>
            <LoginButton />
          </Suspense>
        </div>
      </div>
    </main>
  );
}
