import axios from 'axios';
import { getServerConfig } from '../setup';
import { Event } from 'src/models/events/event';

const serverConfig = getServerConfig();

export const getEvents = async (): Promise<Event[] | string> => {
  try {
    const response = await axios.get(`${serverConfig.api_url}/events`);
    console.log(response.data);
    return response.data;
  } catch (error) {
    console.error('Error retrieving events:', error);
    throw error;
  }
};
