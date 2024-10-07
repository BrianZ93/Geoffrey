import axios from 'axios';
import { getServerConfig } from '../setup';
import { Recipe } from '@/models/cookbook/recipes';

const serverConfig = getServerConfig();

export const addRecipe = async (recipe: Omit<Recipe, 'id'>): Promise<void> => {
  try {
    const response = await axios.post(
      `${serverConfig.api_url}/cookbook/add-recipe`,
      recipe
    );
    console.log('Recipe added successfully:', response.data);
  } catch (error) {
    console.error('Error adding recipe:', error);
    throw error;
  }
};
