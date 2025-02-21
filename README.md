# go-react-web

Have you ever noticed that after listening to Spotify for a while, you get recommended the same songs every time? Whether it be albums, artists, songs, or playlists, Spotify uses [collaborative filtering](https://www.techaheadcorp.com/blog/spotify-recommendation-system/#:~:text=Instead%20of%20just%20individual%20user,deeper%20understanding%20of%20user%20preferences.) to recommend songs to users. This is summarized as Spotify looking for users who like the same songs as you, 
and recommending songs that *they* like. This can be great, but it can also be repetitive. So, my project looks to use Spotify's recommendation 
algorithm but without the user-based side of it, utilizing Spotify's Web API.

The back and frontend for this song recommendation website uses Spotify Web API, Golang, and React. Also other frameworks such as Tailwind, Fiber, etc.

First thing I did was create the React web app with Vite, choosing to write in Typescript + SWC.
Then, I created the go backend directory as well as the go.mod file based on my previous project https://github.com/huynhli/golang-learning.

The backend will run based on Fiber rather than http, so I ran go get github.com/gofiber/fiber/v2 in the terminal, and updated my main.go code.

I then wrote some frontend code/implemented buttons for backend testing, implemented backend routing and api logic to grab genres, and deployed my backend and frontend using Render and Cloudflare Pages.

Implementation for recommendations was done next in a seperate handler file, taking in a list of genres and returning a list of track names and 
their corresponding arists.
- For this step, I tried: get recommendations endpoint, get categories and get playlists based on categories, and even get artists related to artist. All these options are deprecated and have little-to-no return data that I could use.
- Thus, I'm hardcoding playlists for each genre. Ideally the links are stored in some kind of SQL database, but for brevity I devided to write a large switch statement that it chooses songs from. 

Finally, the recommended songs get displayed in the frontend dynamically with scroll panes.

Thanks for checking out my project! Check out my Github profile if you'd like to see more things like this in the future!