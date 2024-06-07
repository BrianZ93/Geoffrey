import axios from 'axios';
import { getServerConfig } from '../setup';
import { Guest } from 'src/models/events/guest';

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
