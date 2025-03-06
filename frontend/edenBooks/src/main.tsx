import { StrictMode } from "react";
import { createRoot } from "react-dom/client";

import { App } from "./App";
import { BrowserRouter } from "react-router";
import { Navbar } from "./components/Navbar";
import { Provider } from "react-redux";
import store from "./Store/store";
import { Footer } from "./components/Footer";
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Provider store={store}>
        <Navbar />
        <App />
        <Footer/>
      </Provider>
    </BrowserRouter>
  </StrictMode>
);
