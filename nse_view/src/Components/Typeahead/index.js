import React, { useState, useEffect } from "react";

import { AsyncTypeahead } from "react-bootstrap-typeahead";

import PropTypes from "prop-types";

require("react-bootstrap-typeahead/css/Typeahead.css");

function Typeahead(props) {
  const [isLoading, setIsLoading] = useState(false);
  const [options, setOptions] = useState([]);

  async function search(query) {
    setIsLoading(true);
    props.loadingCB && props.loadingCB(true);

    try {
      const t = await props.handleSearch(query);
      setOptions(t);
    } catch (error) {
      props.onErrorCB && props.onErrorCB(error);
    }
    props.loadingCB && props.loadingCB(false);
    setIsLoading(false);
  }

  function onChange(selected) {
    props.handleSelected && props.handleSelected(selected);
  }

  return (
    <AsyncTypeahead
      placeholder={props.placeholder}
      isLoading={isLoading}
      onSearch={query => search(query)}
      options={options}
      onChange={selected => onChange(selected)}
      id={props.id || "nse_search_ta"}
    />
  );
}

Typeahead.propTypes = {
  handleSearch: PropTypes.func.isRequired,
  loadingCB: PropTypes.func,
  handleSelected: PropTypes.func,
  placeholder: PropTypes.string
};

export { Typeahead };
