import { App } from './../App';
import { Component, OnInit, Input } from '@angular/core';

@Component({
    selector: 'game-list',
    templateUrl: './game-list.component.html',
    styleUrls: ['./game-list.component.css']
})
export class GameListComponent implements OnInit {

    constructor() { }

    @Input()
    apps!: App[];

    ngOnInit(): void {
    }

    selectApp(event: Event) {
        let app = (event.target as HTMLElement);

        if (app.nodeName !== "BUTTON") {
            app = app.parentElement as HTMLElement;
        }

        document.querySelectorAll(".app-button").forEach((e) => {
            e.classList.remove("current")
        });

        app.classList.toggle("current", true);
    }
}
