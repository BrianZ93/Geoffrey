import axios from 'axios';
import { getServerConfig } from '../setup';
import { Event } from 'src/models/events/event';

const serverConfig = getServerConfig();

export const updateEvent = async (event: Event): Promise<void> => {
  try {
    const response = await axios.put(
      `${serverConfig.api_url}/events/${event.id}`,
      event
    );
    console.log('Event updated successfully:', response.data);
  } catch (error) {
    console.error('Error updating event:', error);
    throw error;
  }
};
