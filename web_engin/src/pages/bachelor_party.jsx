import React, { useEffect, useRef, useState } from "react";
import { fetchItems } from "./aws/aws_config";

const BachelorPartyPage = ({ activePage }) => {
  const iframeRef = useRef(null);
  const [dataToPass, setDataToPass] = useState({
    message: "Hello from React",
    items: [],
  });

  const sendMessage = () => {
    if (iframeRef.current) {
      iframeRef.current.contentWindow?.postMessage(
        dataToPass,
        window.location.origin
      );
    }
  };

  useEffect(() => {
    const fetchData = async () => {
      const items = await fetchItems();
      setDataToPass({ message: "Hello from React", items });
    };

    fetchData();
  }, [activePage]);

  useEffect(() => {
    if (iframeRef.current && iframeRef.current.contentWindow) {
      iframeRef.current.onload = sendMessage;
    }
  }, [activePage, dataToPass]);

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
