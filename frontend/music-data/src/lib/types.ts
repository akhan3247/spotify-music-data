interface UserProfile{
    display_name: string;
    followers: number;
    images: Image[];
}


interface Image {
    url: string;

    width: number;
    height: number;
}

interface Track {
    album: {
        images: Image[];
    };
    name: string;
    artists: {
        name: string;
    }[];
    duration_ms: number;
    external_urls: {
        spotify: string;
    };
}

interface Artist {
    name: string;
    external_urls: {
        spotify: string;
    };
    genres: string[];
    images: Image[];
}

interface TopTracks {
    items: Track[];
}

interface TopArtists {
    items: Artist[];
}

// Export the interfaces
export type {
    Image,
    Track,
    Artist,
    TopTracks,
    TopArtists,
    UserProfile
}; 