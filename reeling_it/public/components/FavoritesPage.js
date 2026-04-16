import { API } from "../services/API.js";
import { CollectionPage } from "./CollectionPage.js";

export class FavoritesPage extends CollectionPage {
    constructor() {
        super(API.getFavorites, "Favorite Movies")
    }
}

customElements.define("favorite-page", FavoritesPage)