import axios from 'axios';
import { getServerConfig } from '../setup';

const serverConfig = getServerConfig();

export const deleteGuest = async (
  eventId: string,
  guestId: string
): Promise<void> => {
  try {
    const response = await axios.delete(
      `${serverConfig.api_url}/events/${eventId}/guests/${guestId}`
    );
    console.log('Guest deleted successfully:', response.data);
  } catch (error) {
    console.error('Error deleting guest:', error);
    throw error;
  }
};
