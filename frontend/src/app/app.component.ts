import { Component, OnInit } from '@angular/core';
import { App } from './App';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
    apps!: App[];
    buttonText: string = "Download";
    currentProject: string = "Monoboy";

    ngOnInit(): void {
        this.apps = [
            {
                name: "test",
                versions: ["0.1.0", "0.2.0", "0.3.0"]
            },
            {
                name: "Monoboy",
                versions: ["0.1.0", "0.2.0", "0.3.0", "0.4.0"]
            },
            {
                name: "Top Down Shooter",
                versions: ["0.1.0", "0.2.0"]
            }
        ]

    }

    getVersions(): string[] {
        let t = this.apps.find((a) => a.name == this.currentProject)?.versions;
        if (t === undefined) {
            return ["ERROR"];
        }
        return t;
    }

    modalClose() {
        let popup = document.getElementById("popup");
        if (popup != undefined) {
            popup.classList.add("modal-closed");
        }
    }

    modalOpen() {
        let popup = document.getElementById("popup");
        if (popup != undefined) {
            popup.classList.remove("modal-closed");
        }
    }

    modalDelete() {
        this.modalClose();
    }

    appInteract() {
        // window.wails
    }
}
