import React, { useEffect, useRef } from "react";
import { ActivePage } from "../App";

interface BachelorPartyPageProps {
  activePage: ActivePage;
}

const dataToPass = { message: "Hello from React", items: [1, 2, 3, 4, 5] };

const BachelorPartyPage: React.FC<BachelorPartyPageProps> = ({
  activePage,
}) => {
  const iframeRef = useRef<HTMLIFrameElement>(null);

  const sendMessage = () => {
    if (iframeRef.current) {
      iframeRef.current.contentWindow?.postMessage(
        dataToPass,
        window.location.origin
      );
    }
  };

  useEffect(() => {
    if (iframeRef.current && iframeRef.current.contentWindow) {
      iframeRef.current.onload = sendMessage;
    }
  }, [activePage]);

  return (
    <iframe
      ref={iframeRef}
      src={`${process.env.PUBLIC_URL}/bachelor_party/public_html/index.html`}
      title="Bachelor Party"
      style={{ width: "100%", height: "100%", border: "none" }}
      onLoad={sendMessage} // Ensure message is sent after iframe is loaded
    />
  );
};

export default BachelorPartyPage;
