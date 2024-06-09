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

export class Address {
  streetNumber: number;
  streetName: string;
  city: string;
  state: string;
  zip: number;

  constructor(
    streetNumber: number,
    streetName: string,
    city: string,
    state: string,
    zip: number
  ) {
    this.streetNumber = streetNumber;
    this.streetName = streetName;
    this.city = city;
    this.state = state;
    this.zip = zip;
  }
}

export class Property {
  id: string;
  address: Address;
  propertyUse: PropertyUse;
  propertyType: PropertyType;
  mortgages: Mortgage[];

  constructor(
    id: string,
    address: Address,
    propertyUse: PropertyUse,
    propertyType: PropertyType,
    mortgages: Mortgage[]
  ) {
    this.id = id;
    this.address = address;
    this.propertyUse = propertyUse;
    this.propertyType = propertyType;
    this.mortgages = mortgages;
  }
}
