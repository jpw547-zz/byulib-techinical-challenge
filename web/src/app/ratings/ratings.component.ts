import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-ratings',
  templateUrl: './ratings.component.html',
  styleUrls: ['./ratings.component.scss']
})
export class RatingsComponent implements OnInit {
  currentItem: Item;

  constructor(private http: HttpClient) {
    this.getItem();
  }

  ngOnInit(): void {
  }

  async getItem() {
    this.http.get("https://api.lib.byu.edu/leaflet/item").toPromise().then((resp) => {
      this.currentItem = resp as Item;
      // console.log(this.currentItem);
    }, (err) => {
      console.error(err);
    })
  }

  getItemVerb() {
    if (this.currentItem) {
      switch (this.currentItem.type) {
        case ItemType.Book:
          return "read";
        case ItemType.Film:
          return "watch"
      }
    }
  }

  getItemIcon() {
    if (this.currentItem) {
      switch (this.currentItem.type) {
        case ItemType.Book:
          return "book";
        case ItemType.Film:
          return "videocam"
      }
    }
  }

  rateItem(rating: boolean) {
    const body = new RatingRequest();
    body.itemId = this.currentItem.id;
    body.rating = rating;
    
    this.http.post("https://api.lib.byu.edu/leaflet/users/jpw547/ratings", body, {
      headers: {
        "Content-Type": "application/json"
      }
    }).toPromise().then((resp) => {
      // console.log(resp);
      this.getItem();
    });
  }
}

enum ItemType {
  Book = "BOOK",
  Film = "FILM"
}

class Item {
  author: string;
  description: string;
  id: string;
  thumbnail: string;
  title: string;
  type: ItemType;
}

class RatingRequest {
  itemId: string;
  rating: boolean;
}
