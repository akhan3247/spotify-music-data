import axios from "axios";
import { 
  TopTracks,
  TopArtists,
  UserProfile} from "@/lib/types"
import Cookies from 'js-cookie';

async function getUserProfile(): Promise<UserProfile> {
  const token = Cookies.get('spotify_token');
  console.log(token);
  const response = await axios.get<UserProfile>('http://localhost:8080/api/profile', {
    headers: {

      Authorization: `Bearer ${token}`,
      
      
    },
  });
  console.log(response.data);
  return response.data;

}

export { getUserProfile };


async function getTopTracks(): Promise<TopTracks>{  
  const token = Cookies.get('spotify_token');
  const response = await axios.get<TopTracks>('http://localhost:8080/api/top-tracks',{
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  console.log(response.data);


  return response.data;
}


export { getTopTracks };

async function getTopArtists(): Promise<TopArtists>{
  const token = Cookies.get('spotify_token');
  const response = await axios.get<TopArtists>('http://localhost:8080/api/top-artists',{
    headers: {
      Authorization: `Bearer ${token}`,
    },
  }); 
  return response.data;
}

export { getTopArtists };


async function getPlaylists() {
  const response = await axios.get('http://localhost:8080/api/playlists');
  return response.data;
}

export { getPlaylists };

async function logout() {
  Cookies.remove('spotify_token');
  await axios.post('http://localhost:8080/api/logout');
}


export { logout };
