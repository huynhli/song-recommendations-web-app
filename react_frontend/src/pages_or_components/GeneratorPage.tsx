import { useEffect, useState } from 'react'
import '../App.css'

export default function GeneratorPage() {
  const [genreList, setGenreList] = useState<string[]>([]);
  const [showGenres, setShowGenres] = useState(false)
  const [genreAPI, setGenreAPI] = useState<string[]>([]);
  const [link, setLink] = useState("");

  const handleAPI = async () => {
    try {
      const response = await fetch(`http://localhost:8080/api/data?link=${link}`)
      const data = await response.json();
      setGenreAPI(data)
      if (genreAPI.length > 1) {
        setShowGenres(!showGenres)
      }
    } catch (error) {
      console.error('Error fetching: ', error)
    }

  }

  const handleLinkChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLink(event.target.value);
  }

  return (
    <div className="bg-blue-700 min-h-[calc(100vh-120px)] flex flex-col justify-center items-center">
      <div>
        <h1 className="text-7xl text-white font-bold pt-20 pb-6">Song Recs</h1>
        <p className='text-white text-center italic pb-6'>For more information, visit the Docs page.</p>
      </div>
      <div className='bg-blue-700 border-b-10 border-blue-700 flex'>
        <input className='bg-white rounded-lg focus:outline-none focus:ring-0 w-100 px-4'
          type="text"
          value={link}
          onChange={handleLinkChange} // Update state as user types
          placeholder="Enter a spotify link here"
        />
        <div className='bg-blue-700 min-w-20'></div>
        <button onClick={handleAPI} className='rounded-lg bg-white px-3 py-2 hover:cursor-pointer'>Click here for genres</button>
        <div></div>
      </div>
      <div className='bg-blue-700 min-h-20 flex justify-center items-center'>
        <p className='text-white'>{genreAPI}</p>
      </div>
      <div className='w-full grid grid-cols-[repeat(auto-fit,_minmax(300px,_1fr))] gap-4 bg-blue-700'>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        <div className="bg-white py-4 text-center">
          <p>{genreAPI}</p>
        </div>
        
      </div>
    </div>
  )
}
