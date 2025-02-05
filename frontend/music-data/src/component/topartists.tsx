'use client'

import { useState, useEffect } from 'react';
import { getTopArtists } from '@/lib/spotify';
import { Artist } from '@/lib/types';

export default function TopArtists() {
    const [topArtists, setTopArtists] = useState<Artist[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    
    useEffect(() => {
        async function fetchTopArtists() {
            try {
                const artistsData = await getTopArtists();
                setTopArtists(artistsData.items.slice(0, 10));
                setLoading(false);
            } catch (err) {
                console.error(err);
                setError('Failed to load top artists');
                setLoading(false);
            }
        }
        fetchTopArtists();
    }, []);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="p-4 bg-gray-900 rounded-lg">
            <h2 className="text-2xl font-bold mb-4">Top Artists</h2>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-10">
                {topArtists.map((artist, index) => (
                    <div 
                        key={index} 

                        className="bg-gray-800 p-4 rounded-lg hover:scale-105 transition-transform gap-10"
                    >
                        <img 
                            src={artist.images[0]?.url || '/default-artist.png'} 

                            alt={artist.name} 
                            className="w-full h-40 object-cover rounded-md" 
                        />
                        <h3 className="mt-2 text-xl text-center font-semibold font-size-20">
                            {artist.name}
                        </h3>
                    </div>

                ))}
            </div>
        </div>
    );
}


