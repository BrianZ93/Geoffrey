import React, { useState } from "react";
import "./App.css";
import BachelorPartyPage from "./pages/bachelor_party";
import HappilyEverAfterPage from "./pages/happily_ever_after_party";

export enum ActivePage {
  BachelorParty,
  HappilyEverAfter,
}

function App() {
  const [activePage, setActivePage] = useState<ActivePage>(
    ActivePage.BachelorParty
  );

  return (
    <div className="App">
      {activePage === ActivePage.BachelorParty && (
        <BachelorPartyPage activePage={activePage} />
      )}
      {activePage === ActivePage.HappilyEverAfter && (
        <HappilyEverAfterPage activePage={activePage} />
      )}
    </div>
  );
}

export default App;