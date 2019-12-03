import React, { useState, useEffect } from "react";

import "bootstrap/dist/css/bootstrap.min.css";

import { Typeahead } from "./Components/Typeahead";

import Header from "./Components/Header";
import ListView from "./Components/ListView";

const data = require("./payload").default;

const initialCorpData = {
  boardMeeting: [],
  corpAction: [],
  corpInfo: {},
  corpAnnouncement: []
};

function App() {
  const [corp, setCorp] = useState("");

  const [isLoadingCorpDetails, setIsLoadingCorpDetails] = useState(false);

  const [corpData, setCorpData] = useState(initialCorpData);

  useEffect(() => {
    if (!corp[0]) {
      setCorpData(initialCorpData);
      setCorp("");
      return;
    }
    setIsLoadingCorpDetails(true);
    App.FetchCorpData(corp[0].split(" ")[0])
      .then(data => {
        setIsLoadingCorpDetails(false);
        setCorpData(data);
      })
      .catch(err => {
        setIsLoadingCorpDetails(false);
        alert("Failed to fetch corp data");
      });

    // TODO fetch data
  }, [corp]);

  return (
    <div>
      <Header>
        <div className="row">
          <div className="col-md-8" style={{ padding: "14px" }}>
            <h2 style={{ display: "inline-block" }}>NSE DEMO</h2>
            <span style={{ marginLeft: "30px" }}>{corp}</span>
          </div>
          <div className="col-md-4" style={{ padding: "14px" }}>
            <Typeahead
              placeholder="Enter Symbol / Company Name"
              handleSearch={q => App.HandleSearch(q)}
              loadingCB={isLoading => null}
              handleSelected={selected => setCorp(selected)}
              key="nse_search"
            />
          </div>
        </div>
      </Header>

      <div className="container-fluid">
        {!corp && (
          <div style={{ textAlign: "center", marginTop: "50px" }}>
            Please select company name
          </div>
        )}

        {isLoadingCorpDetails && (
          <div style={{ textAlign: "center", marginTop: "50px" }}>
            Loading Details of {corp}
          </div>
        )}

        {!isLoadingCorpDetails && corp && (
          <div className="row">
            <div className="col-md-5">
              <div className="shadow outer-b-m">
                <h3>Board Meeting</h3>
                <div className="inner-b-m">
                  <ListView
                    heading="Board Meeting"
                    data={corpData.boardMeeting}
                  />
                </div>
              </div>
            </div>

            <div className="col-md-7">
              <div className="row">
                <div className="rigth-card shadow">
                  <h4>Corporate Action</h4>
                  <ListView
                    heading="Corporate Action"
                    data={corpData.corpAction}
                  />
                </div>

                <div className="shadow rigth-card">
                  <h4>Corporate Info</h4>
                  <ListView heading="corpInfo" data={corpData.corpInfo} />
                </div>

                <div className="shadow rigth-card">
                  <h4>Corporate Announcement</h4>
                  <ListView
                    heading="corpAnnouncement"
                    data={corpData.corpAnnouncement}
                  />
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

// Static function
App.HandleSearch = function HandleSearch(query) {
  return fetch("http://localhost:3333/company?query=" + query)
    .then(async data => {
      const d = await data.json();
      return d.rows1.map(c => c.CompanyValues.replace(/  +/g, " "));
    })
    .catch(err => {
      throw err;
    });
};

App.FetchCorpData = function FetchCorpData(query) {
  return fetch("http://localhost:3333/company/details?symbol=" + query)
    .then(async data => {
      const d = await data.json();
      return d;
    })
    .catch(err => {
      throw err;
    });
};

export default App;
