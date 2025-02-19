# go-react-web
The back and frontend for a song recommendation website using spotify API, Golang, and React. Also other things such as Tailwind, Fiber, etc.

First thing I did was create the React web app with Vite, choosing to write in Typescript + SWC.
Then, I created the go backend directory as well as the go.mod file based on my previous project https://github.com/huynhli/golang-learning.

The backend will run based on Fiber rather than http, so I ran go get github.com/gofiber/fiber/v2 in the terminal, and updated my main.go code.

I then wrote some frontend code/implemented buttons for backend testing, implemented backend routing and api logic to grab genres, and deployed my backend and frontend using Render and Cloudflare Pages.
