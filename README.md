# NSE Website Scrapper

> https://www.nseindia.com/corporates/corporateHome.html----  404 error 

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
    # 1. Clone the repo inside $GOPATH/src folder
    git clone https://github.com/anks333/nse_scrapper.git $GOPATH/src/nse_scrapper

    # 2. Navigate to root dir of the repo
    cd $GOPATH/src/nse_scrapper

    # 3. Build front-end
    cd nse_view
    # install node dependencies
    npm install
    # build react app
    npm run build


    # 4. Navigate back to `nse_scrapper` dir
    cd ..
    # install go dependencies
    go get ./...
    # start server
    go run main.go

    # 5. Browse http://localhost:3333

    ```

> Application runs on port 3333, you can change it by changing the `env`, also need to update the `Dockerfile` and `docker-compose.yml`

## APIs

1. Search Company names

   ```sh
    curl http://localhost:3333/company?query=COMPANY_NAME_OR_SYMBOL
    # e.g. curl http://localhost:3333/company?query=TCS
   ```

   Response

   ```json
   {
     "rows1": [
       {
         "CompanyValues": "INFY  Infosys Limited",
         "CompanyNames": "INFY<br><a style='color:red;'>Infosys</a> Limited"
       },
       {
         "CompanyValues": "HCL-INSYS  HCL Infosystems Limited",
         "CompanyNames": "HCL-INSYS<br>HCL <a style='color:red;'>Infosys</a>tems Limited"
       }
     ],
     "success": "true",
     "results1": 2
   }
   ```

2. Get Company Details
   ```sh
   curl http://localhost:3333/company/details?symbol=COMPANY_SYMBOL
   # e.g. curl http://localhost:3333/company/details?symbol=TCS
   ```
   Response
   ```json
   {
     "boardMeeting": [
       {
         "date": "18-Apr-2017",
         "purpose": "Results/Dividend"
       },
       {
         "date": "20-Feb-2017",
         "purpose": "Buyback"
       },
       {
         "date": "12-Jan-2017",
         "purpose": "Results/Dividend"
       },
       {
         "date": "13-Oct-2016",
         "purpose": "Results/Dividend"
       },
       {
         "date": "14-Jul-2016",
         "purpose": "Results/Dividend"
       },
       {
         "date": "18-Apr-2016",
         "purpose": "Results/Dividend"
       },
       {
         "date": "12-Jan-2016",
         "purpose": "Results/Dividend"
       },
       {
         "date": "13-Oct-2015",
         "purpose": "Results/Dividend"
       },
       {
         "date": "09-Jul-2015",
         "purpose": "Results/Dividend"
       },
       {
         "date": "16-Apr-2015",
         "purpose": "Results/Dividend"
       },
       {
         "date": "15-Jan-2015",
         "purpose": "Results/Dividend"
       },
       {
         "date": "16-Oct-2014",
         "purpose": "Results/Dividend"
       },
       {
         "date": "17-Jul-2014",
         "purpose": "Results/Dividend"
       },
       {
         "date": "16-Apr-2014",
         "purpose": "Results/Dividend"
       },
       {
         "date": "16-Jan-2014",
         "purpose": "Results/Dividend"
       },
       {
         "date": "15-Oct-2013",
         "purpose": "Results/Dividend"
       },
       {
         "date": "18-Jul-2013",
         "purpose": "Results/Dividend"
       },
       {
         "date": "17-Apr-2013",
         "purpose": "Results/Dividend"
       },
       {
         "date": "14-Jan-2013",
         "purpose": "Results/Dividend"
       },
       {
         "date": "19-Oct-2012",
         "purpose": "Miscelleneous"
       },
       {
         "date": "19-Oct-2012",
         "purpose": "Results/Dividend"
       },
       {
         "date": "12-Jul-2012",
         "purpose": "Results/Dividend"
       },
       {
         "date": "23-Apr-2012",
         "purpose": "Results/Dividend"
       },
       {
         "date": "17-Jan-2012",
         "purpose": "Results/Dividend"
       },
       {
         "date": "17-Oct-2011",
         "purpose": "Results/Dividend"
       },
       {
         "date": "14-Jul-2011",
         "purpose": "Results/Dividend"
       },
       {
         "date": "21-Apr-2011",
         "purpose": "Results/Dividend"
       },
       {
         "date": "17-Jan-2011",
         "purpose": "Results/Dividend"
       },
       {
         "date": "21-Oct-2010",
         "purpose": "Results/Dividend"
       }
     ],
     "corpAction": [
       {
         "date": "17-Oct-2019",
         "purpose": "INTERIM DIVIDEND - RS 5 PER SHARE AND SPECIAL DIVIDEND - RS 40 PER SHARE"
       },
       {
         "date": "16-Jul-2019",
         "purpose": "INTERIM DIVIDEND - RS 5 PER SHARE"
       },
       {
         "date": "04-Jun-2019",
         "purpose": "DIVIDEND - RS 18 PER SHARE"
       },
       {
         "date": "17-Jan-2019",
         "purpose": "INTERIM DIVIDEND - RS 4 PER SHARE"
       },
       {
         "date": "23-Oct-2018",
         "purpose": "INT DIV RS 4 PER SH"
       }
     ],
     "corpAnnouncement": [
       "- Press Release  Dec 04, 2019, 12:06",
       "- Press Release  Dec 03, 2019, 14:53",
       "- Press Release  Dec 03, 2019, 12:29",
       "- Press Release  Dec 02, 2019, 16:17",
       "- Analysts/Institutional Investor Meet/Con. Call Updates  Nov 29, 2019, 15:33"
     ],
     "corpInfo": {
       "52 week high/low price": "2296.20/1808.00",
       "Company Name": "Tata Consultancy Services Limited",
       "Constituent Indices": ",NIFTY 50,NIFTY 100,NIFTY SERVICES SECTOR,NIFTY 500,NIFTY 200,NIFTY IT",
       "Date of Listing (NSE)": "25-Aug-2004",
       "Face Value": "1.00",
       "Free Float Market Cap. *": "218381.29(Cr)",
       "ISIN": "INE467B01029",
       "Impact Cost": "0.02 as on Nov-2019",
       "Industry": "COMPUTERS - SOFTWARE",
       "Issued Cap.": "3752384706(shares) as on 04-Dec-2019"
     }
   }
   ```

## Test

> TODO: Write test cases
