import React, { useEffect, useRef, useState, useCallback } from "react";
import { getAllItems, updateGuest } from "./aws-interface";
import { ActivePage } from "../App";

interface BachelorPartyPageProps {
  activePage: ActivePage;
}

interface DataToPass {
  message: string;
  items: Record<string, any>[]; // Adjust this type based on the actual structure of items
}

const BachelorPartyPage: React.FC<BachelorPartyPageProps> = ({
  activePage,
}) => {
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
    const handleCustomEvent = async (event: Event) => {
      const customEvent = event as CustomEvent<{
        email: string;
        note: string;
        rsvp: boolean;
      }>;
      console.log("hi");
      console.log("Received data:", customEvent.detail);

      const guest = dataToPass.items.find(
        (guest) =>
          guest.email.toLowerCase() === customEvent.detail.email.toLowerCase()
      );

      if (guest) {
        try {
          const result = await updateGuest("Events_Guests", guest.id, {
            attending: customEvent.detail.rsvp,
          });
          console.log("Guest updated:", result);
        } catch (error) {
          console.error("Failed to update guest:", error);
        }
      } else {
        console.error("Guest not found");
      }
    };

    window.addEventListener(
      "rsvpSubmitted",
      handleCustomEvent as EventListener
    );

    return () => {
      window.removeEventListener(
        "rsvpSubmitted",
        handleCustomEvent as EventListener
      );
    };
  }, [dataToPass.items]);

  return (
    <iframe
      ref={iframeRef}
      src={`${process.env.PUBLIC_URL}/bachelor_party/public_html/bachelor_party.html`}
      title="Bachelor Party"
      style={{ width: "100%", height: "100%", border: "none" }}
      onLoad={sendMessage} // Ensure message is sent after iframe is loaded
    />
  );
};

export default BachelorPartyPage;
