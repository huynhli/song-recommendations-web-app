# go-react-web

Have you ever noticed that after listening to Spotify for a while, you get recommended the same songs every time? Whether it be albums, artists, songs, or playlists, Spotify uses [collaborative filtering](https://www.techaheadcorp.com/blog/spotify-recommendation-system/#:~:text=Instead%20of%20just%20individual%20user,deeper%20understanding%20of%20user%20preferences.) to recommend songs to users. This is summarized as Spotify looking for users who like the same songs as you, 
and recommending songs that *they* like. This can be great, but it can also be repetitive. So, my project looks to use Spotify's recommendation 
algorithm but without the user-based side of it, utilizing Spotify's Web API.

The back and frontend for this song recommendation website uses Spotify Web API, Golang, and React. Also other frameworks such as Tailwind, Fiber, etc.

This is a React web app with Vite, written in Typescript + SWC.
It is deployed using Render and Cloudflare Pages.

Working with recommendations was hard:
- I tried: the get recommendations endpoint, get categories and get playlists based on categories, and even get artists related to artist. All these options are deprecated and have little-to-no return data that I could use.
- Thus, I hardcoded playlists for each genre. Ideally the links are stored in some kind of database, but for brevity I decided to write a large switch statement that it chooses songs from.

*Note: There are TODOs littered throughout the code for future implementations. Just ignore them lol*

Thanks for checking out my project! Check out my Github profile if you'd like to see more things like this in the future!