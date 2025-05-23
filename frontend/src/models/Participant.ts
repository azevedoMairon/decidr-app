export default class Participant {
  id: string;
  name: string;
  imageUrl: string;
  isEliminated: boolean;
  isNominated: boolean;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.imageUrl = data.image_url;
    this.isEliminated = data.is_eliminated;
    this.isNominated = data.is_nominated;
  }
}
