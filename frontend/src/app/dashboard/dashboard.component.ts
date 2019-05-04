import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  public menuOpened = true;

  constructor(public router: Router) {}

  ngOnInit() {
    if (this.router.url === '/') {
      this.router.navigate(['agreements']);
    }
  }

  toggleMenu() {
    this.menuOpened = !this.menuOpened;
  }
}
