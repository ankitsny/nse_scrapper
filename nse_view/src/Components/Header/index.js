import React from "react";
import { Typeahead } from "../Typeahead";

import PropTypes from "prop-types";

function Header(props) {
  return (
    <div
      className="container-fluid"
      style={{
        backgroundColor: props.bgColor,
        height: props.height,
        boxShadow: "0px 1px 20px gray"
      }}
    >
      {props.children}
    </div>
  );
}

Header.propTypes = {
  bgColor: PropTypes.string.isRequired,
  height: PropTypes.string.isRequired
};

Header.defaultProps = {
  bgColor: "#05aa9e",
  height: "65px"
};

export default Header;
