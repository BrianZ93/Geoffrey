import axios from 'axios';
import { getServerConfig } from '../setup';
import { Guest } from 'src/models/events/guest';

const serverConfig = getServerConfig();

export const modifyGuest = async (
  eventId: string,
  guestId: string,
  guest: Omit<Guest, 'id' | 'eventId'>
): Promise<void> => {
  try {
    const response = await axios.put(
      `${serverConfig.api_url}/events/${eventId}/guests/${guestId}`,
      {
        ...guest,
      }
    );
    console.log('Guest modified successfully:', response.data);
  } catch (error) {
    console.error('Error modifying guest:', error);
    throw error;
  }
};
