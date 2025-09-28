import { useEffect, useState } from 'react'
import '../App.css'
import { getAlbumRec, getArtistRec, getTrackRec } from './TanstackHelper';
import { useQuery } from '@tanstack/react-query'

export default function GeneratorPage() {
  // useSearchParams for query filtering, pagination, anything that edits teh query
  // const [searchParams, setSearchParams] = useSearchParams()
  // const linkQuery = searchParams.get('link')

  const [resourceForQuery, setResourceForQuery] = useState<{ type: string, id: string } | null>(null)
  const [link, setLink] = useState("");
  // TODO impement debouncing
  const handleLinkChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setLink(event.target.value);
  }

  // on page load
  useEffect(() => {
    // simple + stateless
    const url = new URLSearchParams(window.location.search)
    const linkQuery = url.get('link')

    if (linkQuery){
      const decodedLink = decodeURIComponent(linkQuery)
      setLink(decodedLink) // change box text, and state
      sanitizeValidateAndBackendCall(decodedLink) // on page load, use decoded link and not state, bc state updates async
    }
  }, [])

  // on button submit
  const handleSubmit = (e: React.FormEvent) => {
    e?.preventDefault()
    sanitizeValidateAndBackendCall(link)
  }
  
  const sanitizeLink = (link : string) : [string, string | null] => {
    if (!link) return [link, "No string entered"]
    const trimmedLowerLink = link.trim().toLowerCase()
    // must be a spotify link
    if (!trimmedLowerLink.startsWith('https://open.spotify.com/')) {
      return [trimmedLowerLink, 'Not a valid spotify link: Does not start with "https://open.spotify.com/"']
    }
    return [trimmedLowerLink, null]
  }

  const [validationError, setValidationError] = useState<string | null>(null)

  const sanitizeValidateAndBackendCall = (decodedLink : string) => {
    // fully sanitize/validate 
    // TODO add loading anim for checking string vs loading anim for requesting data
    const [sanitizedLink, error] = sanitizeLink(decodedLink)
    if (error) {
      setValidationError(error)
      return
    }
    setValidationError(null)

    console.log('Link passed valididation: ' + sanitizedLink)
    console.log('Requesting info...')
    // backend request, display value/error
    console.log(sanitizedLink?.split('/'))
    const splitSanitizedLink = sanitizedLink?.split('/')
    const id = splitSanitizedLink[4] 
    const linkType = splitSanitizedLink[3]
    if (!["album", "track", "artist"].includes(linkType)){
      return console.log('Spotify link type not supported. Please use a different link.')
    }
    
    setResourceForQuery({type: linkType, id: id})
    // update page info
  }

  const queryFn = resourceForQuery?.type === "album"
    ? () => getAlbumRec(resourceForQuery.id)
    : resourceForQuery?.type === "artist"
    ? () => getArtistRec(resourceForQuery.id)
    : resourceForQuery?.type === "track"
    ? () => getTrackRec(resourceForQuery.id)
    : undefined

  const { data : recs = [], error, isLoading, isFetched } = useQuery({
  queryKey: resourceForQuery ? [resourceForQuery.type, resourceForQuery.id] : [],
  queryFn: resourceForQuery ? queryFn! : () => Promise.resolve([]),
  // enabled: !!resourceForQuery, --> only runs if theres a resourceForQuery, also enables auto refetching
});
  return (
    // TODO make page not scrollable
    <section className='h-[100vh] w-full flex flex-col justify-center items-center'>
        <div className={`h-full w-full flex justify-center items-center flex-2 flex-col`}>
          <form className='max-w-[1600px] mx-[10%] text-black flex flex-row' onSubmit={handleSubmit}>
            <div className='flex justify-center items-center'>
            <input 
              className='flex px-2 mx-2 w-auto max-w-600 min-w-100 bg-white'
              type='text'
              value={link}
              onChange={handleLinkChange}
              placeholder='Enter a spotify link here'
            />
            </div>
            <div className='flex justify-center items-center'>
            <button className='px-2 mx-2 rounded-md   w-30 border-black border-1 bg-gray-300 hover:bg-gray-200 active:bg-gray-100' type="submit">Generate</button>
            </div>
          </form>
          {validationError && (
            <div className='my-2 flex'>{validationError}</div>
          )}
        </div>
        { resourceForQuery && (
          isLoading ? 
          <div className='flex-1 flex justify-center items-center'>currently loading... please wait ...</div> 
          : error ?
          <div className='flex-1 flex justify-center items-center'>{error instanceof Error ? error.message : "An error occurred"}</div>
          : isFetched && recs.length === 0 ?
          <div className='flex-1 flex justify-center items-center'>0 recs</div>
          :
          <div className='flex-3 flex justify-center items-center'>{recs.map((rec:string[], i:number) => (
            <div key={i}>
              {rec}
            </div>
          ))}
          </div>
          
        )}
        

    </section>
  )
}

