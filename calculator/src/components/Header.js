import React, { Component } from "react";
import logo from '../logo.svg'

class Header extends Component {
  render() {
    return (
      <div className="text-center">
        <img
          src={logo}
          width="300"
          className="image-thumbnail"
          style={{ marginTop: "20px" }}
        />
        <h1>Statistic Calculator</h1>
        <hr /> 
      </div>
    );
  }
}

export default Header;