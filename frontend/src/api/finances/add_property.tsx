import axios from 'axios';
import { getServerConfig } from '../setup';
import { Property } from './../../models/finances/property';

const serverConfig = getServerConfig();

export const addProperty = async (
  property: Omit<Property, 'id'>
): Promise<void> => {
  try {
    const response = await axios.post(
      `${serverConfig.api_url}/finances/properties`,
      property
    );
    console.log('Property added successfully:', response.data);
  } catch (error) {
    console.error('Error adding property:', error);
    throw error;
  }
};
