import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route } from "react-router-dom";
import logo from "./logo.svg";
import "./App.css";
import Splash from "./Splash";

import { ThemeProvider } from 'theme-ui';
import rebass from "@rebass/preset";

function App() {
  return (
    <div className="App">
      <ThemeProvider>{Splash}</ThemeProvider>
    </div>
  );
}

export default App;
