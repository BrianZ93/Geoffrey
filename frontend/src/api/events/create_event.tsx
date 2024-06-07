import axios from 'axios';
import { getServerConfig } from '../setup';
import { Event } from 'src/models/events/event';

const serverConfig = getServerConfig();

export const createEvent = async (event: Event): Promise<void> => {
  try {
    const response = await axios.post(
      `${serverConfig.api_url}/create-event`,
      event
    );
    console.log('Event created successfully:', response.data);
  } catch (error) {
    console.error('Error creating event:', error);
    throw error;
  }
};
