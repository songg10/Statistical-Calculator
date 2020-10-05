import React, { Component, Fragment } from "react";
import './App.css';
import Header from "./components/Header";
import Home from "./components/Home";

function App() {
  return (
    <Fragment>
        <Header />
        <Home />
    </Fragment>
  );
}

export default App;
