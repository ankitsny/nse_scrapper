# NSE Website Scrapper

> https://www.nseindia.com/corporates/corporateHome.html

-> Corporate information -> Board meetings in the left menu shows a list of companies on the right side (a table) with a search box on top. When you click on the blue link under "Purpose", it pops up a page with info about that particular board meeting.

## Task:

- Write a standalone HTTP app in Go that presents a search box and a button.
- The search box accepts a stock's tradingsymbol (eg: INFY) and sends an AJAX request to the backend
- On receiving the request, the backend scrapes the NSE page and displays the following results: Board meeting date, Purpose, Details
