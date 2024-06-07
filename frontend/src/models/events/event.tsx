import { Guest } from './guest';

export class Event {
  id: string;
  title: string;
  startTime: string;
  endTime: string;
  guests: Guest[];
  venues: Venue[];

  constructor(
    id: string,
    title: string,
    startTime: string,
    endTime: string,
    guests: Guest[],
    venues: Venue[]
  ) {
    this.id = id;
    this.title = title;
    this.startTime = startTime;
    this.endTime = endTime;
    this.guests = guests;
    this.venues = venues;
  }
}

export class Venue {
  id: string;
  title: string;
  address: string;
  startTime: string;
  endTime: string;
  guests: Guest[];

  constructor(
    id: string,
    title: string,
    address: string,
    startTime: string,
    endTime: string,
    guests: Guest[]
  ) {
    this.id = id;
    this.title = title;
    this.address = address;
    this.startTime = startTime;
    this.endTime = endTime;
    this.guests = guests;
  }
}
