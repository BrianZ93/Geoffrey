import React, { useState } from "react";
import "./App.css";

import BachelorPartyPage from "./pages/bachelor_party";
import HappilyEverAfterPage from "./pages/happily_ever_after_party";

export enum ActivePage {
  BachelorParty,
  HappilyEverAfter,
}

function App() {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [activePage, setActivePage] = useState<ActivePage>(
    ActivePage.BachelorParty
  );

  // GET THIS FROM THE DASHBOARD OR DB TABLE
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const event_id = "76125999-17f0-4f22-b47e-ef4ac16fb1cc"

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
