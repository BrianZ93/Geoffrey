export class Guest {
  id: string;
  name: string;
  email: string;
  phoneNumber: string;
  attending: boolean;
  rsvpReceived: boolean;
  note: string;

  constructor(
    id: string,
    name: string,
    email: string,
    phoneNumber: string,
    attending: boolean,
    rsvpReceived: boolean,
    note: string
  ) {
    this.id = id;
    this.name = name;
    this.email = email;
    this.phoneNumber = phoneNumber;
    this.attending = attending;
    this.rsvpReceived = rsvpReceived;
    this.note = note;
  }
}
