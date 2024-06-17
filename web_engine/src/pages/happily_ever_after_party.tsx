import React from "react";

import { ActivePage } from "../App";

interface HappilyEverAfterPageProps {
  activePage: ActivePage;
}

const HappilyEverAfterPage: React.FC<HappilyEverAfterPageProps> = ({
  activePage,
}) => {
  return (
    <iframe
      src={`${process.env.PUBLIC_URL}/happily_ever_after_party/happily_ever_after_party.html`}
      title="Happily Ever After Party"
      style={{ width: "100%", height: "100%", border: "none" }}
    />
  );
};

export default HappilyEverAfterPage;
