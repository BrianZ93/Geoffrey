import { Country } from '../main';

export enum CookingAppliance {
  None = 0,
  Stove = 1,
  Oven = 2,
  Microwave = 3,
  Toaster = 4,
  ToasterOven = 5,
  Blender = 6,
  FoodProcessor = 7,
  SlowCooker = 8,
  PressureCooker = 9,
  RiceCooker = 10,
  DeepFryer = 11,
  AirFryer = 12,
  OutdoorGrill = 13,
  Griddle = 14,
  SandwichMaker = 15,
  WaffleMaker = 16,
  CoffeeMaker = 17,
  EspressoMachine = 18,
  ElectricKettle = 19,
  SousVide = 20,
  InductionCooktop = 21,
  BreadMaker = 22,
  IceCreamMaker = 23,
  StandMixer = 24,
  HandMixer = 25,
  Juicer = 26,
  HotPlate = 27,
  Rotisserie = 28,
  Smokers = 29,
  DutchOven = 30,
  InstantPot = 31,
  Steamer = 32,
  FonduePot = 33,
  PopcornMaker = 34,
  EggCooker = 35,
  Unknown = 36,
}

export enum CookingAction {
  Other = 0,
  Stir = 1,
  Mix = 2,
  Whisk = 3,
  Beat = 4,
  Fold = 5,
  Knead = 6,
  Chop = 7,
  Slice = 8,
  Dice = 9,
  Mince = 10,
  Grate = 11,
  Shred = 12,
  Peel = 13,
  Julienne = 14,
  Sear = 15,
  Saute = 16,
  Fry = 17,
  DeepFry = 18,
  Boil = 19,
  Simmer = 20,
  Blanch = 21,
  Poach = 22,
  Steam = 23,
  Roast = 24,
  Bake = 25,
  Broil = 26,
  Grill = 27,
  Braise = 28,
  Stew = 29,
  Marinate = 30,
  Glaze = 31,
  Baste = 32,
  Caramelize = 33,
  Deglaze = 34,
  Puree = 35,
  Blend = 36,
  Chill = 37,
  Freeze = 38,
  Thaw = 39,
  Garnish = 40,
  Measure = 41,
  Weigh = 42,
  Sift = 43,
  Strain = 44,
  Season = 45,
  Brine = 46,
  Smoke = 47,
  Cure = 48,
  Ferment = 49,
  Proof = 50,
  Skim = 51,
  Toast = 52,
  Crush = 53,
  Grind = 54,
  Melt = 55,
  Reheat = 56,
  Rest = 57,
  Trim = 58,
}

export enum RecipeStepStage {
  Preparation = 0,
  Cooking = 1,
  Resting = 2,
}

export class FlavorProfile {
  sweet: number;
  salty: number;
  spicy: number;
  sour: number;
  savory: number;

  constructor(
    sweet: number,
    salty: number,
    spicy: number,
    sour: number,
    savory: number
  ) {
    this.sweet = this.validateValue(sweet);
    this.salty = this.validateValue(salty);
    this.spicy = this.validateValue(spicy);
    this.sour = this.validateValue(sour);
    this.savory = this.validateValue(savory);
  }

  private validateValue(value: number): number {
    if (value < 0) {
      return 0;
    } else if (value > 5) {
      return 5;
    } else {
      return value;
    }
  }
}

export class Ingredient {
  id: string;
  name: string;

  constructor(id: string, name: string) {
    this.id = id;
    this.name = this.capitalizeWords(name);
  }

  private capitalizeWords(value: string): string {
    return value.replace(/\b\w/g, (char) => char.toUpperCase());
  }
}

export class RecipeStep {
  id: string;
  name: string;
  applianceUsed: CookingAppliance;
  cookingAction: CookingAction;
  timeSeconds: number;
  prepTimeSeconds: number;
  ingredients: Ingredient[];
  description: string;
  recipeStage: RecipeStepStage;

  constructor(
    id: string,
    name: string,
    appliancedUsed: CookingAppliance,
    cookingAction: CookingAction,
    timeSeconds: number,
    prepTimeSeconds: number,
    ingredients: Ingredient[],
    description: string,
    recipeStage: RecipeStepStage
  ) {
    this.id = id;
    this.name = name;
    this.applianceUsed = appliancedUsed;
    this.cookingAction = cookingAction;
    this.timeSeconds = timeSeconds;
    this.prepTimeSeconds = prepTimeSeconds;
    this.ingredients = ingredients;
    this.description = description;
    this.recipeStage = recipeStage;
  }
}

export class Recipe {
  id: string;
  name: string;
  steps: RecipeStep[];
  totalTimeSeconds: number;
  totalTimePrepSeconds: number;
  flavorProfile: FlavorProfile;
  country: Country;
  ingredients: Ingredient[];
  note?: string;

  constructor(
    id: string,
    name: string,
    steps: RecipeStep[],
    totalTimeSeconds: number,
    totalTimePrepSeconds: number,
    flavorProfile: FlavorProfile,
    country: Country,
    ingredients: Ingredient[],
    note?: string
  ) {
    this.id = id;
    this.name = name;
    this.steps = steps;
    this.totalTimeSeconds = totalTimeSeconds;
    this.totalTimePrepSeconds = totalTimePrepSeconds;
    this.flavorProfile = flavorProfile;
    this.country = country;
    this.ingredients = ingredients;
    this.note = note;
  }
}
