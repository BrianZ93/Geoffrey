import { Property } from './property';

export class Assets {
  id: string;
  properties: Property[];

  constructor(id: string, properties: Property[]) {
    this.id = id;
    this.properties = properties;
  }
}

export class Portfolio {
  id: string;
  assets: Assets[];

  constructor(id: string, assets: Assets[]) {
    this.id = id;
    this.assets = assets;
  }
}
