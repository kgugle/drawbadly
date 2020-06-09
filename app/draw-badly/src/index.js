import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route } from "react-router-dom";
import "./index.css";
import * as serviceWorker from "./serviceWorker";

import { ThemeProvider } from 'emotion-theming'
import theme from './theme'
import Splash from "./Splash";
import NewGame from "./NewGame";
import JoinGame from "./JoinGame";

ReactDOM.render(
  <ThemeProvider theme={theme}>
    <BrowserRouter>
      <Route exact path="/" component={Splash} />
      <Route exact path="/new" component={NewGame} />
      <Route exact path="/join" component={JoinGame} />
    </BrowserRouter>
  </ThemeProvider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
