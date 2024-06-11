import { Mortgage } from './debt';

export enum PropertyUse {
  Residential = 0,
  Investment = 0,
}

export enum PropertyType {
  House = 0,
  Townhouse = 1,
  Condo = 2,
  MultiUnit = 3,
  Commercial = 4,
}

export enum LeaseType {
  MonthToMonth = 0,
  FixedPeriod = 1,
}

export class Income {
  id: string;
  description: string;
  memo: string;
  amount: number;
  date: string;

  constructor(
    id: string,
    description: string,
    memo: string,
    amount: number,
    date: string
  ) {
    this.id = id;
    this.description = description;
    this.memo = memo;
    this.amount = amount;
    this.date = date;
  }
}

export class Expense {
  id: string;
  description: string;
  memo: string;
  amount: number;
  date: string;

  constructor(
    id: string,
    description: string,
    memo: string,
    amount: number,
    date: string
  ) {
    this.id = id;
    this.description = description;
    this.memo = memo;
    this.amount = amount;
    this.date = date;
  }
}

export class Address {
  id: string;
  streetNumber: number;
  streetName: string;
  city: string;
  state: string;
  zip: number;

  constructor(
    id: string,
    streetNumber: number,
    streetName: string,
    city: string,
    state: string,
    zip: number
  ) {
    this.id = id;
    this.streetNumber = streetNumber;
    this.streetName = streetName;
    this.city = city;
    this.state = state;
    this.zip = zip;
  }
}

export class Tenant {
  id: string;
  firstName: string;
  lastName: string;
  address: Address;
  leaseType: LeaseType;
  leaseStartDate: string;
  leaseEndDate: string;
  rent: number;
  securityDeposit: number;
  otherIncome: number;

  constructor(
    id: string,
    firstName: string,
    lastName: string,
    address: Address,
    leaseType: LeaseType,
    leaseStartDate: string,
    leaseEndDate: string,
    rent: number,
    securityDeposit: number,
    otherIncome: number
  ) {
    this.id = id;
    this.firstName = firstName;
    this.lastName = lastName;
    this.address = address;
    this.leaseType = leaseType;
    this.leaseStartDate = leaseStartDate;
    this.leaseEndDate = leaseEndDate;
    this.rent = rent;
    this.securityDeposit = securityDeposit;
    this.otherIncome = otherIncome;
  }
}

export class Property {
  id: string;
  purchaseDate: string;
  purchasePrice: number;
  address: Address;
  propertyUse: PropertyUse;
  propertyType: PropertyType;
  mortgages: Mortgage[];
  tenants: Tenant[];
  income: Income[];
  expenses: Expense[];

  constructor(
    id: string,
    purchaseDate: string,
    purchasePrice: number,
    address: Address,
    propertyUse: PropertyUse,
    propertyType: PropertyType,
    mortgages: Mortgage[],
    tenants: Tenant[],
    income: Income[],
    expenses: Expense[]
  ) {
    this.id = id;
    this.purchaseDate = purchaseDate;
    this.purchasePrice = purchasePrice;
    this.address = address;
    this.propertyUse = propertyUse;
    this.propertyType = propertyType;
    this.mortgages = mortgages;
    this.tenants = tenants;
    this.income = income;
    this.expenses = expenses;
  }
}
