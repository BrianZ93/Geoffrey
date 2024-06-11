import { Portfolio } from '../../models/finances/portfolio';

export const getMockPortfolio = () => {
  const mock_portfolio = new Portfolio('0', []);
  return mock_portfolio;
};
