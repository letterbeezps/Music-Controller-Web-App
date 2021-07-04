import React, { Component } from "react";
import HomePage from "./components/HomePage";

import "./App.css";

class App extends Component {
  // eslint-disable-next-line no-useless-constructor
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="App">
        <HomePage />
      </div>
    );
  }
}

export default App;
