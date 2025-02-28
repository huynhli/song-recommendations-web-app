import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [genreList, setGenreList] = useState<string[]>([]);
  const [genreAPI, setGenreAPI] = useState<string[]>([]);
  const [link, setLink] = useState("");

  const handleClick = () => {
    // Fetch token from backend using fetch API
    fetch("https://song-recommendations-web-app-7jyz.onrender.com/data") //sends GET to backend at that link
      .then((response) => {
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }
        return response.json();
      })
      .then((data) => {
        setGenreList(data);
      })
      .catch((error) => {
        console.error("Error fetching token:", error);
      });
  };

  const handleAPI = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/data?link=${link}`)
      const data = await response.json();
      setGenreAPI(data)
    } catch (error) {
      console.error('Error fetching: ', error)
    }

  }

  const handleLinkChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLink(event.target.value);
  }



  return (
    <div className="bg-blue-300 min-h-screen flex justify-center items-center">
      <div>
        <h1>Vite + React</h1>
        <h1 className="text-3xl font-bold underline">Hello yall</h1>
        <div className="card">
          <input
            type="text"
            value={link}
            onChange={handleLinkChange} // Update state as user types
            placeholder="Enter link"
          />
          <button onClick={handleAPI}>Click here for genres</button>
          <p>{genreAPI} here is API</p>
          <button onClick={handleClick}>Click here for genres</button>
          <p>{genreList} here</p>

        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </div>
    </div>
  )
}

export default App
