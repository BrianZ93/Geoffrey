import axios from 'axios';
import * as yaml from 'yaml';
import { getServerConfig } from '../setup';
import { Guest } from './../../models/events/guest';

const serverConfig = getServerConfig();

export const addGuest = async (
  eventId: string,
  guest: Omit<Guest, 'id' | 'eventId'>
): Promise<void> => {
  try {
    const response = await axios.post(
      `${serverConfig.api_url}/events/${eventId}/guests`,
      {
        ...guest,
        eventId,
      }
    );
    console.log('Guest added successfully:', response.data);
  } catch (error) {
    console.error('Error adding guest:', error);
    throw error;
  }
};

// Function to import guests from files
export function addGuestsFromYAML(file: File): Promise<Guest[]> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = (event) => {
      if (event.target?.result) {
        const fileContent = event.target.result as string;
        console.log('Adding guests from file content');

        try {
          const parsed = yaml.parse(fileContent);
          const guests: Guest[] = parsed.guests.map(
            (guest: Guest) =>
              new Guest(
                '0',
                guest.name,
                guest.email,
                guest.phoneNumber,
                guest.attending,
                guest.rsvpReceived,
                guest.note
              )
          );

          resolve(guests);
        } catch (error) {
          reject(`Error parsing YAML: ${error}`);
        }
      } else {
        reject('File content is empty');
      }
    };
    reader.onerror = () => reject('Error reading file');
    reader.readAsText(file);
  });
}
