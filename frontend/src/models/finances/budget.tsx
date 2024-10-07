export enum Frequency {
  Daily = 0,
  Weekly = 1,
  Monthly = 2,
  Quarterly = 3,
  SemiAnnually = 4,
  Annually = 5,
}

export class Expense {
  id: string;
  description: string;
  amount: number;
  frequency: Frequency;
  category?: string;
  expiration?: Date;

  constructor(
    id: string,
    description: string,
    amount: number,
    frequency: Frequency,
    category: string,
    expiration: Date
  ) {
    this.id = id;
    this.description = description;
    this.amount = amount;
    this.frequency = frequency;
    this.category = category;
    this.expiration = expiration;
  }
}

export class Income {
  id: string;
  description: string;
  amount: number;
  frequency: Frequency;
  category?: string;
  expiration?: Date;

  constructor(
    id: string,
    description: string,
    amount: number,
    frequency: Frequency,
    category: string,
    expiration: Date
  ) {
    this.id = id;
    this.description = description;
    this.amount = amount;
    this.frequency = frequency;
    this.category = category;
    this.expiration = expiration;
  }
}

export class Budget {
  id: string;
  categories: string[];
  income: Income[];
  expenses: Expense[];

  constructor(
    id: string,
    categories: string[],
    income: Income[],
    expenses: Expense[]
  ) {
    this.id = id;
    this.categories = categories;
    this.income = income;
    this.expenses = expenses;
  }
}
