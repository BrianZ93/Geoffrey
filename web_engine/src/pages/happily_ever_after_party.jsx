import React from "react";

const HappilyEverAfterPage = ({ activePage }) => {
  return (
    <iframe
      src={`${process.env.PUBLIC_URL}/happily_ever_after_party/index.html`}
      title="Happily Ever After Party"
      style={{ width: "100%", height: "100%", border: "none" }}
    />
  );
};

export default HappilyEverAfterPage;
