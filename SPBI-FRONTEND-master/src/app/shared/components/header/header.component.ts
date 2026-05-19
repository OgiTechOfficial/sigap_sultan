import { Component, OnInit } from '@angular/core';
import { INav } from './nav.interface';
import { MENU } from './nav.resource';
import { openCloseMobileMenu, openCloseProfileMenuAnimation, rotateAnimation } from './header.animation';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss'],
  animations: [openCloseProfileMenuAnimation, rotateAnimation, openCloseMobileMenu]
})
export class HeaderComponent implements OnInit {
  menus: INav[] = MENU;
  hideProfileMenu: boolean = true;
  hideNavbar: boolean = true;

  constructor() { }

  ngOnInit(): void {
  }

  signout() {
    localStorage.clear();
  }

}
