'use client';

import { getUserProfile } from "@/lib/spotify";
import { useEffect, useState } from "react";
import type { UserProfile } from "@/lib/types";

export default function UserProfile() {

  const [userProfile, setUserProfile] = useState<UserProfile | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function loadProfile() {
      try {
        const profile = await getUserProfile();
        setUserProfile(profile);
      } catch (err) {
        console.error(err);
        setError("Failed to load user profile");
      }
    }
    loadProfile();
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!userProfile) {
    return <div>Loading...</div>;
  }

  return (
    <div className="max-w-4xl mx-auto p-6">
      <div className="flex flex-col md:flex-row items-center gap-8">
        {/* Left Container - Profile Picture */}
        <div className="flex-shrink-0">
          <div className="w-[200px] h-[200px] md:w-[240px] md:h-[240px] rounded-full overflow-hidden ring-2 ring-gray-700">
            {userProfile.images && userProfile.images.length > 0 ? (
              <img
                src={userProfile.images[0].url}
                alt={userProfile.display_name}
                className="w-[200px] h-[200px] md:w-[240px] md:h-[240px] rounded-full object-cover"
              />
            ) : (
              <div className="w-full h-full bg-gray-700 flex items-center justify-center">
                <span className="text-2xl text-gray-400">?</span>
              </div>
            )}
          </div>
        </div>

        {/* Right Container - User Info */}
        <div className="flex flex-col text-center md:text-left">
          <h2 className="text-6xl md:text-7xl font-bold text-white tracking-tight">
            {userProfile.display_name}
          </h2>
          <div className="flex items-center justify-center md:justify-start gap-2 text-[#9B9B9B] mt-2">
            <span className="text-lg md:text-xl">{userProfile.followers} followers</span>
          </div>
        </div>
      </div>
    </div>
  );
}
