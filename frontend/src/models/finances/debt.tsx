export enum DebtType {
  Installment = 0,
  Revolving = 1,
  Balloon = 2,
}

export class PrincipleOnlyPayment {
  id: string;
  date: string;
  amount: number;

  constructor(id: string, date: string, amount: number) {
    this.id = id;
    this.date = date;
    this.amount = amount;
  }
}

export class EscrowBreakdown {
  id: string;
  propertyTaxes: number;
  interest: number;
  associationDues: number;

  constructor(
    id: string,
    propertyTaxes: number,
    interest: number,
    associationDues: number
  ) {
    this.id = id;
    this.propertyTaxes = propertyTaxes;
    this.interest = interest;
    this.associationDues = associationDues;
  }
}

export class Mortgage {
  id: string;
  debtType: DebtType.Installment;
  originationDate: string;
  firstPaymentDate: string;
  maturity: string;
  interestRate: number;
  lienholderPosition: number;
  servicer: string;
  address: string;
  principleOnlyPayments: PrincipleOnlyPayment[];
  term: number;
  escrow: boolean;
  escrowBreakdown: EscrowBreakdown;

  constructor(
    id: string,
    originationDate: string,
    firstPaymentDate: string,
    maturity: string,
    interestRate: number,
    lienholderPosition: number,
    servicer: string,
    address: string,
    principleOnlyPayments: PrincipleOnlyPayment[],
    term: number,
    escrow: boolean,
    escrowBreakdown: EscrowBreakdown
  ) {
    this.id = id;
    this.debtType = DebtType.Installment;
    this.originationDate = originationDate;
    this.firstPaymentDate = firstPaymentDate;
    this.maturity = maturity;
    this.interestRate = interestRate;
    this.lienholderPosition = lienholderPosition;
    this.servicer = servicer;
    this.address = address;
    this.principleOnlyPayments = principleOnlyPayments;
    this.term = term;
    this.escrow = escrow;
    this.escrowBreakdown = escrowBreakdown;
  }
}
