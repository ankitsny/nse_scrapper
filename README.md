# NSE Website Scrapper

> https://www.nseindia.com/corporates/corporateHome.html

-> Corporate information -> Board meetings in the left menu shows a list of companies on the right side (a table) with a search box on top. When you click on the blue link under "Purpose", it pops up a page with info about that particular board meeting.

## Task:

- Write a standalone HTTP app in Go that presents a search box and a button.
- The search box accepts a stock's tradingsymbol (eg: INFY) and sends an AJAX request to the backend
- On receiving the request, the backend scrapes the NSE page and displays the following results:

1. Board meeting date,
2. Purpose,
3. Details

## Prerequisite

1. golang environment
2. node.js (to build the front-end, react app)

   OR

3. Docker

## To start the app

1. Using Docker


    ```sh
    # Pull from docker hub
    docker run -p 3333:3333 ankso3/nse_scrapper

    # OR
    docker build --rm -f "Dockerfile" -t nse_scrapper:latest .
    docker run -p 3333:3333 nse_scrapper

    # OR
    docker-compose -f "docker-compose.yml" up -d --build
    ```

2. Without Docker


    ```sh
    #1. navigate to `nse_scrapper/nse_view` folder

    # install node dependencies
    npm install

    # build react app
    npm run build


    #2. Navigate back to `nse_scrapper` dir

    # install go dependencies
    go get ./...

    # start server
    go run main.go
    ```

> Application runs on port 3000, you can change it by changing the `env`, also need to update the `Dockerfile` and `docker-compose.yml`
