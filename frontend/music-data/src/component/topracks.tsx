'use client'

import { useState, useEffect } from 'react';
import { getTopTracks } from '@/lib/spotify';
import { Track } from '@/lib/types';

/**
 * Formats the duration (in ms) to a string of the format "m:ss"
 */
function formatDuration(durationMs: number): string {
    const minutes = Math.floor(durationMs / 60000);
    const seconds = Math.floor((durationMs % 60000) / 1000);
    return `${minutes}:${seconds < 10 ? '0' + seconds : seconds}`;
}

export default function TopTracks() {
    const [topTracks, setTopTracks] = useState<Track[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        async function fetchTopTracks() {
            try {
                const tracksData = await getTopTracks();
                // Only take the first 10 tracks
                setTopTracks(tracksData.items.slice(0, 10));
                setLoading(false);
            } catch (err) {
                console.error(err);
                setError('Failed to load top tracks');
                setLoading(false);
            }
        }
        fetchTopTracks();
    }, []);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="top-tracks-container">
            <ul className="top-tracks-list" style={{ listStyleType: 'none', padding: 0, margin: 0 }}>
                {topTracks.map((track, index) => (
                    <li key={index} className="track-item">
                        <div
                            className="track-row"
                            style={{
                                display: 'flex',
                                alignItems: 'center',
                                padding: '10px 0',
                                borderBottom: '1px solid #ddd'
                            }}
                        >
                            {/* Album Cover */}
                            <img
                                src={track.album.images[0]?.url}
                                alt={track.name}
                                style={{
                                    width: '50px',
                                    height: '50px',
                                    objectFit: 'cover',
                                    marginRight: '15px'
                                }}

                            />

                            {/* Track Info */}
                            <div className="track-details" style={{ flexGrow: 1 }}>
                                <div className="track-name" style={{ fontWeight: 'bold' }}>
                                    {track.name}
                                </div>
                                <div className="track-artist" style={{ color: '#555', fontSize: '0.9em' }}>
                                    {track.artists.map((artist) => artist.name).join(', ')}
                                </div>
                            </div>

                            {/* Track Duration */}
                            <div className="track-duration" style={{ marginLeft: 'auto', fontSize: '0.9em', color: '#777' }}>
                                {formatDuration(track.duration_ms)}
                            </div>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
}
        