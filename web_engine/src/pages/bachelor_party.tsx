import React, { useEffect, useRef, useState, useCallback } from "react";
import { getAllItems } from "./aws-interface";
import { ActivePage } from "../App";

interface BachelorPartyPageProps {
  activePage: ActivePage;
}

interface DataToPass {
  message: string;
  items: Record<string, any>[]; // Adjust this type based on the actual structure of items
}

const BachelorPartyPage: React.FC<BachelorPartyPageProps> = ({ activePage }) => {
  const iframeRef = useRef<HTMLIFrameElement | null>(null);
  const [dataToPass, setDataToPass] = useState<DataToPass>({
    message: "Hello from React",
    items: [],
  });

  const sendMessage = useCallback(() => {
    if (iframeRef.current) {
      iframeRef.current.contentWindow?.postMessage(
        dataToPass,
        window.location.origin
      );
    }
  }, [dataToPass]);

  useEffect(() => {
    const fetchData = async () => {
      const items = await getAllItems("Events_Guests");
      setDataToPass({ message: "Hello from React", items: items || [] });
    };

    fetchData();
  }, [activePage]);

  useEffect(() => {
    if (iframeRef.current && iframeRef.current.contentWindow) {
      iframeRef.current.onload = sendMessage;
    }
  }, [activePage, sendMessage]);

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
