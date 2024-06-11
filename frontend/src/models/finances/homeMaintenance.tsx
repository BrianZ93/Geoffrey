export enum FrequencyType {
  Daily = 0,
  Weekly = 1,
  Monthly = 2,
  Quarterly = 3,
  SemiAnnually = 4,
  Annually = 5,
}

export enum ContractorType {
  GeneralContractor = 0,
  Electrician = 1,
  Plumber = 2,
  Carpenter = 3,
  Painter = 4,
  Roofer = 5,
  HVACTechnician = 6,
  Landscaper = 7,
  Mason = 8,
  FlooringSpecialist = 9,
  WindowInstaller = 10,
  SidingInstaller = 11,
  InsulationInstaller = 12,
  PestControl = 13,
  PoolInstaller = 14,
  FenceInstaller = 15,
  GutterInstaller = 16,
  ConcreteSpecialist = 17,
  SolarPanelInstaller = 18,
  SecuritySystemInstaller = 19,
  Myself = 20,
}

export class Contractor {
  id: string;
  name: string;
  phoneNumber: number;
  email: string;
  contractorType: ContractorType;

  constructor(
    id: string,
    name: string,
    phoneNumber: number,
    email: string,
    contractorType: ContractorType
  ) {
    this.id = id;
    this.name = name;
    this.phoneNumber = phoneNumber;
    this.email = email;
    this.contractorType = contractorType;
  }
}

export class MaintenanceItem {
  id: string;
  description: string;
  frequency: number;
  frequencyType: FrequencyType;
  contractor: Contractor;
  startDate: string;
  endDate?: string;

  constructor(
    id: string,
    description: string,
    frequency: number,
    frequencyType: FrequencyType,
    contractor: Contractor,
    startDate: string,
    endDate: string
  ) {
    this.id = id;
    this.description = description;
    this.frequency = frequency;
    this.frequencyType = frequencyType;
    this.contractor = contractor;
    this.startDate = startDate;
    this.endDate = endDate;
  }
}
