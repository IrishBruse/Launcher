import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
    tests!: any[];
    buttonText: string = "Download";

    ngOnInit(): void {
        this.tests = [{ name: "test" }, { name: "b" }, { name: "c" }];
    }
}
